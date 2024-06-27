package worker

import (
	"context"
	"crypto/ecdsa"
	"github.com/pkg/errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/patrickmn/go-cache"

	"github.com/carv-protocol/verifier/api/gasless"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/pkg/contract"
	"github.com/carv-protocol/verifier/pkg/tools"
)

const (
	// Maximum number of blocks per log query
	maxBlocksPerQuery = 100
)

var (
	nonRetriableErrors = []error{core.ErrInsufficientFunds, core.ErrInsufficientFundsForTransfer}
)

type Chain struct {
	cf                         *conf.Bootstrap
	logger                     *log.Helper
	ethClient                  *ethclient.Client
	protocolServiceContractObj *contract.Contract
	verifierAddress            common.Address
	verifierPrivKey            *ecdsa.PrivateKey
	nodeInf                    nodeInf
	stopChan                   chan struct{}
	verifyResultMap            *tools.SafeMap
	confirmVrfNodeChan         chan confirmVrfNodesInfo
	errChan                    chan error
	latestBlockNumber          int64
	cache                      *cache.Cache

	gaslessClient gasless.GasslessHTTPClient
}

type confirmVrfNodesInfo struct {
	nodeId    uint32
	requestId *big.Int
}

func NewChain(
	ctx context.Context,
	bootstrap *conf.Bootstrap,
	logger *log.Helper,
) (*Chain, error) {
	ethClient, err := ethclient.DialContext(ctx, bootstrap.Chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "new eth client error, rpc url: %s", bootstrap.Chain.RpcUrl)
	}

	// init contract object
	protocolServiceContractObj, err := contract.NewContract(common.HexToAddress(bootstrap.Contract.Addr), ethClient)
	if err != nil {
		return nil, errors.Wrapf(err, "NewContract error")
	}

	privateKey := key_manager.Inst().PrivateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	verifierAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// get nodeInfor
	nodeInfos, err := protocolServiceContractObj.NodeInfos(&bind.CallOpts{}, verifierAddress)

	cacheIns := cache.New(5*time.Minute, 10*time.Minute)

	// init gasless client
	httpClient, err := http.NewClient(ctx, http.WithEndpoint(bootstrap.GaslessService.Url))
	if err != nil {
		return nil, errors.Wrapf(err, "new http client error, url: %s", bootstrap.GaslessService.Url)
	}
	gaslessClient := gasless.NewGasslessHTTPClient(httpClient)

	return &Chain{
		cf:                         bootstrap,
		logger:                     logger,
		ethClient:                  ethClient,
		protocolServiceContractObj: protocolServiceContractObj,
		verifierAddress:            verifierAddress,
		verifierPrivKey:            privateKey,
		nodeInf: nodeInf{
			nodeId:         nodeInfos.Id,
			nodeListIndex:  nodeInfos.ListIndex,
			rewardClaimer:  nodeInfos.Claimer,
			commissionRate: nodeInfos.CommissionRate,
		},
		stopChan:           make(chan struct{}),
		confirmVrfNodeChan: make(chan confirmVrfNodesInfo),
		verifyResultMap:    tools.NewSafeMap(),
		errChan:            make(chan error),
		cache:              cacheIns,
		gaslessClient:      gaslessClient,
	}, nil
}

func (c *Chain) Start(ctx context.Context) error {
	// Retrieve the latest block number if it has not been specified
	if c.latestBlockNumber == 0 {
		blockNumber, err := c.ethClient.BlockNumber(ctx)
		if err != nil {
			return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chain.ChainName)
		}

		c.latestBlockNumber = int64(blockNumber)
	}
	// Apply the offset to begin verification starting from the latest unverified attestation
	c.latestBlockNumber -= c.cf.Chain.GetOffsetBlock()
	//c.latestBlockNumber = 58206220
	c.logger.WithContext(ctx).Infof("chain [%s] startBlockNumber: %d", c.cf.Chain.ChainName, c.latestBlockNumber)

	for {
		// check again
		time.Sleep(2 * time.Second)
		nodeInfos, err := c.protocolServiceContractObj.NodeInfos(&bind.CallOpts{}, c.verifierAddress)
		if err != nil {
			c.logger.Errorf("NodeInfos error: %s", err.Error())
		}
		c.logger.WithContext(ctx).Infof("nodeInfos: %v", nodeInfos)
		if c.beforeScanEvent(ctx, nodeInfos.Claimer, nodeInfos.CommissionRate, true) {
			break
		}
	}

	c.logger.WithContext(ctx).Infof("chain [%s] beforeScanEvent success", c.cf.Chain.ChainName)
	// Monitor and verify on-chain TEE attestations
	go c.queryAndVerify(ctx)

	// submit verify result to gasless service
	go c.submitVerifyResult(ctx)

	// Monitor errors arising from querying and posting processes
	go c.monitorErrors(ctx)

	return nil
}

func (c *Chain) Stop(ctx context.Context) error {
	c.logger.WithContext(ctx).Infof("chain [%s] stopping", c.cf.Chain.ChainName)

	close(c.stopChan)

	c.logger.WithContext(ctx).Infof("chain [%s] stopped", c.cf.Chain.ChainName)

	return nil
}

// Query on-chain events and verify the TEE attestation
func (c *Chain) queryAndVerify(ctx context.Context) {
	eventQueryTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-eventQueryTicker.C:
			go func() {
				if err := c.queryChain(ctx); err != nil {
					c.errChan <- err
				}
			}()
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain queryAndVerify [%s] stopping", c.cf.Chain.ChainName)
			return
		}
	}
}

func (c *Chain) submitVerifyResult(ctx context.Context) {
	for {
		select {
		case cvn := <-c.confirmVrfNodeChan:
			go func(cvn confirmVrfNodesInfo) {
				// get attestations
				attestationsObj, _ := c.verifyResultMap.Get(cvn.requestId.String())
				// Send ConfirmVrfNodes
				resultEnum := 0
				if !attestationsObj.(verifyResult).result {
					resultEnum = 1
				}
				_, err := NodeReportVerificationBatchByGaslessService(ctx, c, attestationsObj.(verifyResult).attestationID, uint8(resultEnum), c.nodeInf.nodeListIndex)
				if err != nil {
					return
				}

				// c.errChan <- err
				c.logger.WithContext(ctx).Infof("attestation Sending: %s", attestationsObj.(verifyResult).attestationID)
				defer c.verifyResultMap.Delete(cvn.requestId.String())
			}(cvn)

		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain [%s] submitVerifyResult stopping", c.cf.Chain.ChainName)
			return
		}
	}
}

// Monitor errors during both querying and posting processes
func (c *Chain) monitorErrors(ctx context.Context) {
	for {
		select {
		case err := <-c.errChan:
			for _, nonRetriableErr := range nonRetriableErrors {
				// Switch from error.Is to message-based checking for enhanced flexibility.
				if errorContainsMessage(nonRetriableErr, err) {
					// TODO: Implement graceful shutdown handling.
					panic(err.Error())
				}
			}
			c.logger.WithContext(ctx).Errorf("mointor error: %s", err.Error())
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain [%s] batchPost stopping", c.cf.Chain.ChainName)
			return
		}
	}
}

func (c *Chain) SendAttestationTrx(ctx context.Context, attestationIds [][32]byte, results []bool) (string, error) {
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.verifierAddress)
	if err != nil {
		return "", errors.Wrap(err, "eth client get Nonce error")
	}

	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "eth client get SuggestGasPrice error")
	}

	chainID, err := c.ethClient.ChainID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "eth client get ChainID error")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.verifierPrivKey, chainID)
	if err != nil {
		return "", errors.Wrap(err, "bind NewKeyedTransactorWithChainID error")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	//tx, err := c.contractObj.VerifyAttestationBatch(auth, attestationIds, results)
	//if err != nil {
	//	return "", errors.Wrap(err, "contract VerifyAttestationBatch error")
	//}
	//
	//txHash := tx.Hash().Hex()

	//c.logger.WithContext(ctx).Infof("tx hash: %s", txHash)

	return "", nil
}

func (c *Chain) queryChain(ctx context.Context) error {
	startBlockNumber := c.latestBlockNumber
	endBlockNumber, err := c.ethClient.BlockNumber(ctx)
	if err != nil {
		return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chain.ChainName)
	}

	// If there is no new blocks
	if startBlockNumber >= int64(endBlockNumber) {
		return nil
	}

	// Limit the maximum number of blocks that can be queried per run
	if startBlockNumber+maxBlocksPerQuery < int64(endBlockNumber) {
		endBlockNumber = uint64(startBlockNumber + maxBlocksPerQuery)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(startBlockNumber),
		ToBlock:   big.NewInt(int64(endBlockNumber)),
		Addresses: []common.Address{
			common.HexToAddress(c.cf.Contract.Addr),
		},
		Topics: [][]common.Hash{
			{
				common.HexToHash(c.cf.Contract.Topic1),
				common.HexToHash(c.cf.Contract.Topic2),
			},
		},
	}

	cLogs, err := c.ethClient.FilterLogs(ctx, query)
	if err != nil {
		return errors.Wrap(err, "client FilterLogs error")
	}

	var logFilter LogFilter
	for _, cLog := range cLogs {
		// teeReportVerificationBatch Event
		if strings.EqualFold(cLog.Topics[0].String(), c.cf.Contract.Topic1) {
			err := logFilter.NodeReportVerificationBatchLogFilter(ctx, c, cLog)
			if err != nil {
				log.Errorf("NodeReportVerificationBatchLogFilter error: %s", err.Error())
				continue
			}
		}

		if strings.EqualFold(cLog.Topics[0].String(), c.cf.Contract.Topic2) {
			err := logFilter.ConfirmVrfNodesLogFilter(ctx, c, cLog)
			if err != nil {
				log.Errorf("NodeReportVerificationLogFilter error: %s", err.Error())
				continue
			}

		}

	}

	c.logger.WithContext(ctx).Infof(
		"chain [%s] query: start block %d, end block %d",
		c.cf.Chain.ChainName,
		startBlockNumber,
		endBlockNumber,
	)
	c.latestBlockNumber = int64(endBlockNumber) + 1

	return nil
}

// Batch process multiple attestations across a specified range of blocks
func (c *Chain) sendBatchResult(ctx context.Context, resultList []verifyResult) error {
	attestationIds := make([][32]byte, len(resultList))
	verifyResults := make([]bool, len(resultList))
	for i, r := range resultList {
		attestationIds[i] = r.attestationID
		verifyResults[i] = r.result
	}
	_, err := c.SendAttestationTrx(ctx, attestationIds, verifyResults)
	if err != nil {
		return errors.Wrapf(err, "chain [%s] send batch tx failed", c.cf.Chain.ChainName)
	}

	return nil
}

func (c *Chain) beforeScanEvent(ctx context.Context, rewardClaimer common.Address, commissionRate uint32, isGasless bool) bool {
	nodeAddress := c.verifierAddress
	// Get the current timestamp
	timestamp := big.NewInt(time.Now().Unix())
	expiredTime := new(big.Int).Add(timestamp, big.NewInt(c.cf.Signature.ExpiredTime))
	// Check if the node is online
	isOnline, err := c.checkIsOnline(nodeAddress)
	if err != nil {
		c.logger.Errorf("checkIsOnline error: %s", err.Error())
		return false
	}
	if !isOnline {
		c.logger.Errorf("node [%s] is offline, waiting online", nodeAddress.Hex())
		replaceNodeReq := &gasless.ExplorerReplacedNodeRequest{
			VerifierAddr: nodeAddress.Hex(),
		}
		replacedNodeResp, err := c.gaslessClient.ExplorerReplacedNode(context.Background(), replaceNodeReq)
		if err != nil {
			return false
		}
		replaceNode := common.HexToAddress(replacedNodeResp.Data.ReplacedAddr)
		// Send Transaction
		if isGasless {
			_, err = NodeEnterByGaslessService(context.Background(), c, replaceNode, expiredTime)
			if err != nil {
				c.logger.Errorf("NodeEnterByGaslessService error: %s", err.Error())
				return false
			}
		} else {
			auth, _ := c.GetTransactionConfig(context.Background())
			enter, err := NodeEnter(ctx, c, auth, replaceNode)
			if err != nil {
				c.logger.WithContext(ctx).Errorf("NodeEnter error: %s", err.Error())
			}
			c.logger.WithContext(ctx).Infof("NodeEnter success: %s", enter.Hex())
		}
		return false
	} else {
		if strings.ToLower(c.cf.Wallet.RewardClaimerAddr) != strings.ToLower(rewardClaimer.Hex()) {
			// Send Transaction
			_, err2 := UpdateNodeRewardClaimerByGaslessService(ctx, c, common.HexToAddress(c.cf.Wallet.RewardClaimerAddr), expiredTime)
			if err2 != nil {
				return false
			}
		} else {
			c.logger.WithContext(ctx).Infof("Update rewardClaimer: %s success.", rewardClaimer.Hex())
		}
		if int(c.cf.Wallet.CommissionRate) != int(commissionRate) {
			// Send Transaction
			_, err := UpdateNodeCommissionRateByGaslessService(context.Background(), c, uint32(c.cf.Wallet.CommissionRate), expiredTime)
			if err != nil {
				return false
			}
		} else {
			c.logger.WithContext(ctx).Infof("Update commissionRate: %d success.", commissionRate)
		}
		if strings.ToLower(c.cf.Wallet.RewardClaimerAddr) == strings.ToLower(rewardClaimer.Hex()) && int(c.cf.Wallet.CommissionRate) == int(commissionRate) {
			return true
		}
	}

	return false
}

// Check node is or not online
func (c *Chain) checkIsOnline(nodeAddress common.Address) (bool, error) {
	nodeInfos, err := c.protocolServiceContractObj.NodeInfos(&bind.CallOpts{}, nodeAddress)
	if err != nil {
		return false, err
	}
	return nodeInfos.Active, nil
}

func (c *Chain) GetTransactionConfig(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.verifierAddress)
	if err != nil {
		return nil, errors.Wrap(err, "eth client get Nonce error")
	}
	// prod gasPrice should be changed
	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "eth client get SuggestGasPrice error")
	}

	chainID, err := c.ethClient.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "eth client get ChainID error")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(c.verifierPrivKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind NewKeyedTransactorWithChainID error")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	return auth, nil
}

// TODO Get Information from gasless service
func (c *Chain) isNodeCountLimitation() (bool, error) {

	return true, nil
}
