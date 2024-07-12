package worker

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
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
	"github.com/pkg/errors"

	"github.com/carv-protocol/verifier/api/gasless"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/pkg/contract"
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
	verifyResultList           []verifyResult
	confirmVrfNodeChan         chan confirmVrfNodesInfo
	exitNodeChan               chan struct{}
	errChan                    chan error
	latestBlockNumber          int64
	cache                      *cache.Cache

	gaslessClient gasless.GasslessHTTPClient
}

type confirmVrfNodesInfo struct {
	nodeId       uint32
	vrfNodeIndex uint32
	requestId    *big.Int
	deadline     *big.Int
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
	httpClient, err := http.NewClient(ctx, http.WithTimeout(30*time.Second), http.WithEndpoint(bootstrap.GaslessService.Url))
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
		verifyResultList:   make([]verifyResult, 0),
		exitNodeChan:       make(chan struct{}),
		errChan:            make(chan error),
		latestBlockNumber:  bootstrap.Chain.StartBlock,
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
		c.logger.WithContext(ctx).Infof("chain [%s] latestBlockNumber: %d", c.cf.Chain.ChainName, c.latestBlockNumber)
	}
	// Apply the offset to begin verification starting from the latest unverified attestation
	c.latestBlockNumber -= c.cf.Chain.GetOffsetBlock()
	//c.latestBlockNumber = 58784070
	//c.latestBlockNumber = 59109506
	c.logger.WithContext(ctx).Infof("chain [%s] startBlockNumber: %d", c.cf.Chain.ChainName, c.latestBlockNumber)

	nodeInfos, err := c.protocolServiceContractObj.NodeInfos(&bind.CallOpts{}, c.verifierAddress)
	if err != nil {
		c.logger.Errorf("NodeInfos error: %s", err.Error())
	}
	c.logger.WithContext(ctx).Infof("nodeInfos: %v", nodeInfos)
	if c.beforeScanEvent(ctx, nodeInfos.Id, nodeInfos.Claimer, nodeInfos.CommissionRate) {
		// Monitor and verify on-chain TEE attestations
		go c.queryAndVerify(ctx)
		//	// submit verify result to gasless service
		go c.submitVerifyResult(ctx)
		//	// Monitor errors arising from querying and posting processes
		go c.monitorErrors(ctx)
	}
	// monitor exit
	go c.monitorExit(ctx)

	return nil
}

func (c *Chain) Stop(ctx context.Context) error {
	c.logger.WithContext(ctx).Infof("chain [%s] stopping", c.cf.Chain.ChainName)

	close(c.stopChan)

	c.logger.WithContext(ctx).Infof("chain [%s] stopped", c.cf.Chain.ChainName)
	// exit node by gasless
	// Get the current timestamp
	timestamp := big.NewInt(time.Now().Unix())
	expiredTime := new(big.Int).Add(timestamp, big.NewInt(c.cf.Signature.ExpiredTime))
	exitRes, err := NodeExitByGaslessService(ctx, c, expiredTime)
	if err != nil {
		return errors.Wrap(err, "exit node by gasless error")
	}
	c.logger.WithContext(ctx).Infof("exitRes: %v", exitRes)

	return nil
}

// Query on-chain events and verify the TEE attestation
func (c *Chain) queryAndVerify(ctx context.Context) {
	queryTicker := c.cf.Chain.QueryTicker
	eventQueryTicker := time.NewTicker(time.Duration(queryTicker) * time.Second)

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
			//delay x seconds
			if c.cf.Chain.ReportDelay == 0 {
				c.cf.Chain.ReportDelay = 30
			}
			time.Sleep(time.Duration(c.cf.Chain.ReportDelay) * time.Second)
			go func(cvn confirmVrfNodesInfo) {
				if cvn.deadline.Cmp(big.NewInt(time.Now().Unix())) == -1 {
					return
				}
				for i := 0; i < len(c.verifyResultList); i++ {
					if c.verifyResultList[i].requestId.Cmp(cvn.requestId) == 0 && !c.verifyResultList[i].isReported {
						resultEnum := 0
						if !c.verifyResultList[i].result {
							resultEnum = 1
						}
						reportRes, err := NodeReportVerificationBatchByGaslessService(ctx, c, c.verifyResultList[i].attestationID, uint8(resultEnum), cvn.vrfNodeIndex)
						if err != nil {
							continue
						}
						if reportRes {
							c.verifyResultList[i].isReported = true
						}
					}
				}
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
	if startBlockNumber+c.cf.Chain.MaxBlockPerQuery < int64(endBlockNumber) {
		endBlockNumber = uint64(startBlockNumber + c.cf.Chain.MaxBlockPerQuery)
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
				common.HexToHash(c.cf.Contract.Topic3),
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
			err1 := logFilter.NodeReportVerificationBatchLogFilter(ctx, c, cLog)
			if err1 != nil {
				log.Errorf("NodeReportVerificationBatchLogFilter error: %s", err1.Error())
				continue
			}
		}

		if strings.EqualFold(cLog.Topics[0].String(), c.cf.Contract.Topic2) {
			err2 := logFilter.ConfirmVrfNodesLogFilter(ctx, c, cLog)
			if err2 != nil {
				log.Errorf("NodeReportVerificationLogFilter error: %s", err2.Error())
				continue
			}

		}

		// NodeClear (exit)
		if strings.EqualFold(cLog.Topics[0].String(), c.cf.Contract.Topic3) {
			err3 := logFilter.NodeClearLogFilter(ctx, c, cLog)
			if err3 != nil {
				log.Errorf("NodeClearLogFilter error: %s", err3.Error())
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

func (c *Chain) beforeScanEvent(ctx context.Context, nodeID uint32, rewardClaimer common.Address, commissionRate uint32) bool {
	nodeAddress := c.verifierAddress
	// Get the current timestamp
	timestamp := big.NewInt(time.Now().Unix())
	expiredTime := new(big.Int).Add(timestamp, big.NewInt(c.cf.Signature.ExpiredTime))

	// Not the first time setup, update config if needed.
	if nodeID != 0 {
		if res := c.updateNodeConfigIfNeeded(ctx, rewardClaimer, commissionRate, expiredTime); !res {
			c.logger.WithContext(ctx).Error("Update node config failed")
			return res
		}
	}

	// Check if the node is online
	isOnline, err := c.checkIsOnline(nodeAddress)
	if err != nil {
		c.logger.Errorf("checkIsOnline error: %s", err.Error())
		return false
	}

	// if node is not online or first time register on chain
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
		enterRes, err := NodeEnterByGaslessService(context.Background(), c, replaceNode, expiredTime)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("NodeEnterByGaslessService error: %s", err.Error())
		}

		if !enterRes {
			c.logger.WithContext(ctx).Error("enterRes is false")
			return enterRes
		}

		// Need to set commission and rewards if it is first time setup
		if nodeID == 0 {
			// Fix: add delay to make sure gasless server detects the node online event
			time.Sleep(time.Duration(c.cf.Chain.ReportDelay) * time.Second)

			res := c.updateNodeConfigIfNeeded(ctx, rewardClaimer, commissionRate, expiredTime)
			if !res {
				c.logger.WithContext(ctx).Error("Update node config failed")
			}
			return res
		}

	}
	return true
}

func (c *Chain) updateNodeConfigIfNeeded(ctx context.Context, rewardClaimer common.Address, commissionRate uint32, expiredTime *big.Int) bool {
	c.logger.WithContext(ctx).Infof(
		"Update node config if needed. On-chain claimer: %s, config claimer: %s; on-chain commission: %d, config commission: %d",
		rewardClaimer.Hex(),
		c.cf.Wallet.RewardClaimerAddr,
		commissionRate,
		c.cf.Wallet.CommissionRate,
	)
	// gas model: reward claimer is current node address
	if !c.cf.Chain.EnableGasMode {
		if strings.ToLower(c.cf.Wallet.RewardClaimerAddr) != strings.ToLower(rewardClaimer.Hex()) {
			c.logger.WithContext(ctx).Infof("update reward claimer address to %s", c.cf.Wallet.RewardClaimerAddr)
			// Send Transaction
			updateRewardClaimerRes, err2 := UpdateNodeRewardClaimerByGaslessService(ctx, c, common.HexToAddress(c.cf.Wallet.RewardClaimerAddr), expiredTime)
			if err2 != nil {
				c.logger.WithContext(ctx).Errorf("UpdateNodeRewardClaimerByGaslessService error: %s", err2.Error())
			}
			return updateRewardClaimerRes
		}
	}

	if int(c.cf.Wallet.CommissionRate) != int(commissionRate) {
		c.logger.WithContext(ctx).Infof("update commission rate to %d", c.cf.Wallet.CommissionRate)
		// Send Transaction
		updateNodeCommissionRateRes, err := UpdateNodeCommissionRateByGaslessService(context.Background(), c, uint32(c.cf.Wallet.CommissionRate), expiredTime)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("UpdateNodeCommissionRateByGaslessService error: %s", err.Error())
		}
		return updateNodeCommissionRateRes
	}

	return true
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
	//TODO Set gasLimit

	return auth, nil
}

func (c *Chain) monitorExit(ctx context.Context) {
	for {
		select {
		case <-c.exitNodeChan:
			c.logger.WithContext(ctx).Infof("Beacause of Node's weight is less than other node, Be replaced!")
			os.Exit(0)
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain [%s] monitorExit stopping", c.cf.Chain.ChainName)
			return
		}
	}
}
