package worker

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/pkg/contract"
)

const (
	// Buffer size for channel storing verification results
	verifyResultChannelCapacity = 10

	// Maximum number of blocks per log query
	maxBlocksPerQuery = 100
)

type Chain struct {
	cf                 *conf.Bootstrap
	logger             *log.Helper
	cAbis              []abi.ABI
	ethClients         []*ethclient.Client
	contractObjs       []*contract.Contract
	verifierAddressArr []common.Address
	verifierPrivKeyArr []*ecdsa.PrivateKey

	stopChans          []chan struct{}
	verifyResultChans  []chan verifyResult
	errChans           []chan error
	latestBlockNumbers []int64
}

type verifyResult struct {
	attestationID [32]byte
	result        bool
}

func NewChain(
	ctx context.Context,
	bootstrap *conf.Bootstrap,
	logger *log.Helper,
) (*Chain, error) {
	abiArr := make([]abi.ABI, 0)
	ethclientArr := make([]*ethclient.Client, 0)
	contractObjs := make([]*contract.Contract, 0)
	privateKeyArr := make([]*ecdsa.PrivateKey, 0)
	verifierAddressArr := make([]common.Address, 0)
	verifyResultChans := make([]chan verifyResult, len(bootstrap.Chains))
	stopChans := make([]chan struct{}, 0)
	errChans := make([]chan error, len(bootstrap.Chains))
	latestBlockNumbers := make([]int64, len(bootstrap.Chains))
	for i := 0; i < len(bootstrap.Chains); i++ {

		abiFile, err := os.ReadFile(bootstrap.Contracts[i].Abi)
		if err != nil {
			return nil, err
		}

		cAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
		if err != nil {
			return nil, errors.Wrap(err, "abi json error")
		}
		abiArr = append(abiArr, cAbi)

		ethClient, err := ethclient.DialContext(ctx, bootstrap.Chains[i].RpcUrl)
		if err != nil {
			return nil, errors.Wrapf(err, "new eth client error, rpc url: %s", bootstrap.Chains[i].RpcUrl)
		}
		ethclientArr = append(ethclientArr, ethClient)

		contractObj, err := contract.NewContract(common.HexToAddress(bootstrap.Contracts[i].Addr), ethClient)
		if err != nil {
			return nil, errors.Wrapf(err, "NewContract error")
		}
		contractObjs = append(contractObjs, contractObj)

		privateKeyBytes, err := Sm4Decrypt(bootstrap, i)
		if err != nil {
			return nil, errors.Wrap(err, "pk decrypt error")
		}
		privateKey, err := crypto.HexToECDSA(string(privateKeyBytes))
		if err != nil {
			return nil, errors.Wrap(err, "pk sm4 HexToECDSA error")
		}
		privateKeyArr = append(privateKeyArr, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		verifierAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		verifierAddressArr = append(verifierAddressArr, verifierAddress)

		stopChans = append(stopChans, make(chan struct{}))
	}
	return &Chain{
		cf:                 bootstrap,
		logger:             logger,
		cAbis:              abiArr,
		ethClients:         ethclientArr,
		contractObjs:       contractObjs,
		verifierAddressArr: verifierAddressArr,
		verifierPrivKeyArr: privateKeyArr,

		stopChans:          stopChans,
		verifyResultChans:  verifyResultChans,
		errChans:           errChans,
		latestBlockNumbers: latestBlockNumbers,
	}, nil
}

func (c *Chain) Start(ctx context.Context, chainId int) error {
	c.latestBlockNumbers[chainId] = c.cf.Chains[chainId].StartBlock

	// Retrieve the latest block number if it has not been specified
	if c.latestBlockNumbers[chainId] == 0 {
		blockNumber, err := c.ethClients[chainId].BlockNumber(ctx)
		if err != nil {
			return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chains[chainId].ChainName)
		}

		c.latestBlockNumbers[chainId] = int64(blockNumber)
	}

	c.logger.WithContext(ctx).Infof("chain [%s] startBlockNumber: %d", c.cf.Chains[chainId].ChainName, c.latestBlockNumbers[chainId])

	// Monitor and verify on-chain TEE attestations
	go c.queryAndVerify(ctx, chainId)

	// Batch post verification results to the blockchain
	go c.batchPost(ctx, chainId)

	// Monitor errors arising from querying and posting processes
	go c.monitorErrors(ctx, chainId)

	return nil
}

func (c *Chain) Stop(ctx context.Context, chainId int) error {
	c.logger.WithContext(ctx).Infof("chain [%s] stopping", c.cf.Chains[chainId].ChainName)

	close(c.stopChans[chainId])

	// todo stop

	c.logger.WithContext(ctx).Infof("chain [%s] stopped", c.cf.Chains[chainId].ChainName)

	return nil
}

// Query on-chain events and verify the TEE attestation
func (c *Chain) queryAndVerify(ctx context.Context, chainId int) {
	eventQueryTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-eventQueryTicker.C:
			go func() {
				if err := c.queryChain(ctx, chainId); err != nil {
					c.errChans[chainId] <- err
				}
			}()
		case <-c.stopChans[chainId]:
			c.logger.WithContext(ctx).Infof("chain queryAndVerify [%s] stopping", c.cf.Chains[chainId].ChainName)
			return
		}
	}
}

// Batch verify results and post them on-chain
func (c *Chain) batchPost(ctx context.Context, chainId int) {
	batchVerifyResultTicker := time.NewTicker(1 * time.Second)
	resultList := make([]verifyResult, 0)

	for {
		select {
		case r := <-c.verifyResultChans[chainId]:
			resultList = append(resultList, r)
		case <-batchVerifyResultTicker.C:
			if len(resultList) == 0 {
				continue
			}

			go func(rl []verifyResult) {
				if err := c.sendBatchResult(ctx, rl, chainId); err != nil {
					c.errChans[chainId] <- err
				}
			}(resultList[:])

			resultList = make([]verifyResult, 0)
		case <-c.stopChans[chainId]:
			c.logger.WithContext(ctx).Infof("chain batchPost [%s] stopping", c.cf.Chains[chainId].ChainName)
			return
		}
	}
}

// Monitor errors during both querying and posting processes
func (c *Chain) monitorErrors(ctx context.Context, chainId int) {
	for {
		select {
		case err := <-c.errChans[chainId]:
			c.logger.WithContext(ctx).Errorf("mointor error: %s", err.Error())
		case <-c.stopChans[chainId]:
			c.logger.WithContext(ctx).Infof("chain batchPost [%s] stopping", c.cf.Chains[chainId].ChainName)
			return
		}
	}
}

func (c *Chain) SendAttestationTrx(ctx context.Context, attestationIds [][32]byte, results []bool, chainId int) (string, error) {
	nonce, err := c.ethClients[chainId].PendingNonceAt(ctx, c.verifierAddressArr[chainId])
	if err != nil {
		return "", errors.Wrap(err, "eth client get Nonce error")
	}

	gasPrice, err := c.ethClients[chainId].SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "eth client get SuggestGasPrice error")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.verifierPrivKeyArr[chainId], big.NewInt(int64(chainId)))
	if err != nil {
		return "", errors.Wrap(err, "bind NewKeyedTransactorWithChainID error")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	tx, err := c.contractObjs[chainId].VerifyAttestationBatch(auth, attestationIds, results)
	if err != nil {
		return "", errors.Wrap(err, "contract VerifyAttestationBatch error")
	}

	txHash := tx.Hash().Hex()

	c.logger.WithContext(ctx).Infof("tx hash: %s", txHash)

	return txHash, nil
}

func (c *Chain) queryChain(ctx context.Context, chainId int) error {
	startBlockNumber := c.latestBlockNumbers[chainId]
	endBlockNumber, err := c.ethClients[chainId].BlockNumber(ctx)
	if err != nil {
		return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chains[chainId].ChainName)
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
			common.HexToAddress(c.cf.Contracts[chainId].Addr),
		},
		Topics: [][]common.Hash{
			{
				common.HexToHash(c.cf.Contracts[chainId].Topic),
			},
		},
	}

	cLogs, err := c.ethClients[chainId].FilterLogs(ctx, query)
	if err != nil {
		return errors.Wrap(err, "client FilterLogs error")
	}

	for _, cLog := range cLogs {
		unpackedData, err := c.contractObjs[chainId].ParseReportTeeAttestation(cLog)
		if err != nil {
			return errors.Wrap(err, "contract ParseReportTeeAttestation error")
		}

		logInfo := LogInfo{
			BlockNumber:        cLog.BlockNumber,
			ContractAddress:    cLog.Address,
			TxHash:             cLog.TxHash,
			TxIndex:            unpackedData.Raw.TxIndex,
			TeeAddress:         unpackedData.TeeAddress,
			CampaignId:         unpackedData.CampaignId,
			AttestationIdBytes: unpackedData.AttestationId,
			AttestationIdStr:   hex.EncodeToString(unpackedData.AttestationId[:]),
			Attestation:        unpackedData.Attestation,
		}

		// Verify attestation
		result, err := verifyAttestation(c, unpackedData.Attestation, chainId)
		if err != nil {
			// If attestation is unable to be parsed and verified, this attestation should be ignored by all verifiers
			c.logger.WithContext(ctx).Error(
				"verify failed, attestation id: %s, error: %s",
				hex.EncodeToString(unpackedData.AttestationId[:]),
				err.Error(),
			)
			continue
		}

		c.verifyResultChans[chainId] <- verifyResult{
			attestationID: unpackedData.AttestationId,
			result:        result,
		}

		c.logger.WithContext(ctx).Infof("logInfo: %+v", logInfo)
	}

	c.logger.WithContext(ctx).Infof("start block %d, end block %d", startBlockNumber, endBlockNumber)
	c.latestBlockNumbers[chainId] = int64(endBlockNumber) + 1

	return nil
}

// Batch process multiple attestations across a specified range of blocks
func (c *Chain) sendBatchResult(ctx context.Context, resultList []verifyResult, chainId int) error {
	attestationIds := make([][32]byte, len(resultList))
	verifyResults := make([]bool, len(resultList))
	for i, r := range resultList {
		attestationIds[i] = r.attestationID
		verifyResults[i] = r.result
	}
	_, err := c.SendAttestationTrx(ctx, attestationIds, verifyResults, chainId)
	if err != nil {
		return err
	}

	return nil
}
