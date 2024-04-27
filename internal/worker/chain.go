package worker

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	"math/big"
	"strings"
	"time"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/pkg/contract"
	"github.com/carv-protocol/verifier/pkg/dcap"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/pkg/errors"
)

const (
	// Buffer size for channel storing verification results
	verifyResultChannelCapacity = 10

	// Maximum number of blocks per log query
	maxBlocksPerQuery = 100
)

type Chain struct {
	cf              *conf.Bootstrap
	logger          *log.Helper
	cAbi            abi.ABI
	ethClient       *ethclient.Client
	contractObj     *contract.Contract
	verifierAddress common.Address
	verifierPrivKey *ecdsa.PrivateKey

	stopChan          chan struct{}
	verifyResultChan  chan verifyResult
	errChan           chan error
	latestBlockNumber int64
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
	ethClient, err := ethclient.DialContext(ctx, bootstrap.Chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "new eth client error, rpc url: %s", bootstrap.Chain.RpcUrl)
	}

	contractObj, err := contract.NewContract(common.HexToAddress(bootstrap.Contract.Addr), ethClient)
	if err != nil {
		return nil, errors.Wrapf(err, "NewContract error")
	}

	cAbi, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		return nil, errors.Wrap(err, "abi json error")
	}

	privateKey := key_manager.Inst().PrivateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	verifierAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Chain{
		cf:              bootstrap,
		logger:          logger,
		cAbi:            cAbi,
		ethClient:       ethClient,
		contractObj:     contractObj,
		verifierAddress: verifierAddress,
		verifierPrivKey: privateKey,

		stopChan:         make(chan struct{}),
		verifyResultChan: make(chan verifyResult, verifyResultChannelCapacity),
		errChan:          make(chan error),
	}, nil
}

func (c *Chain) Start(ctx context.Context) error {
	c.latestBlockNumber = c.cf.Chain.StartBlock

	// Retrieve the latest block number if it has not been specified
	if c.latestBlockNumber == 0 {
		blockNumber, err := c.ethClient.BlockNumber(ctx)
		if err != nil {
			return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chain.ChainName)
		}

		c.latestBlockNumber = int64(blockNumber)
	}

	c.logger.WithContext(ctx).Infof("chain [%s] startBlockNumber: %d", c.cf.Chain.ChainName, c.latestBlockNumber)

	// Monitor and verify on-chain TEE attestations
	go c.queryAndVerify(ctx)

	// Batch post verification results to the blockchain
	go c.batchPost(ctx)

	// Monitor errors arising from querying and posting processes
	go c.monitorErrors(ctx)

	return nil
}

func (c *Chain) Stop(ctx context.Context) error {
	c.logger.WithContext(ctx).Infof("chain [%s] stopping", c.cf.Chain.ChainName)

	close(c.stopChan)

	// todo stop

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

// Batch verify results and post them on-chain
func (c *Chain) batchPost(ctx context.Context) {
	batchVerifyResultTicker := time.NewTicker(1 * time.Second)
	resultList := make([]verifyResult, 0)

	for {
		select {
		case r := <-c.verifyResultChan:
			resultList = append(resultList, r)
		case <-batchVerifyResultTicker.C:
			if len(resultList) == 0 {
				continue
			}

			go func(rl []verifyResult) {
				if err := c.sendBatchResult(ctx, rl); err != nil {
					c.errChan <- err
				}
			}(resultList[:])

			resultList = make([]verifyResult, 0)
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain batchPost [%s] stopping", c.cf.Chain.ChainName)
			return
		}
	}
}

// Monitor errors during both querying and posting processes
func (c *Chain) monitorErrors(ctx context.Context) {
	for {
		select {
		case err := <-c.errChan:
			c.logger.WithContext(ctx).Errorf("mointor error: %s", err.Error())
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain batchPost [%s] stopping", c.cf.Chain.ChainName)
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

	tx, err := c.contractObj.VerifyAttestationBatch(auth, attestationIds, results)
	if err != nil {
		return "", errors.Wrap(err, "contract VerifyAttestationBatch error")
	}

	txHash := tx.Hash().Hex()

	c.logger.WithContext(ctx).Infof("tx hash: %s", txHash)

	return txHash, nil
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
				common.HexToHash(c.cf.Contract.Topic),
			},
		},
	}

	cLogs, err := c.ethClient.FilterLogs(ctx, query)
	if err != nil {
		return errors.Wrap(err, "client FilterLogs error")
	}

	for _, cLog := range cLogs {
		unpackedData, err := c.contractObj.ParseReportTeeAttestation(cLog)
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
		result, err := dcap.VerifyAttestation(unpackedData.Attestation)
		if err != nil {
			// If attestation is unable to be parsed and verified, this attestation should be ignored by all verifiers
			c.logger.WithContext(ctx).Error(
				"verify failed, attestation id: %s, error: %s",
				hex.EncodeToString(unpackedData.AttestationId[:]),
				err.Error(),
			)
			continue
		}

		c.verifyResultChan <- verifyResult{
			attestationID: unpackedData.AttestationId,
			result:        result,
		}

		c.logger.WithContext(ctx).Infof("logInfo: %+v", logInfo)
	}

	c.logger.WithContext(ctx).Infof("start block %d, end block %d", startBlockNumber, endBlockNumber)
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
		return err
	}

	return nil
}
