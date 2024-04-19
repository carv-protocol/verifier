package worker

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"
	"strings"
	"sync"
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
	"github.com/carv-protocol/verifier/internal/data"
	"github.com/carv-protocol/verifier/pkg/contract"
)

const (
	defaultBlockNumberChanLength = 10
)

type Chain struct {
	cf                            *conf.Bootstrap
	data                          *data.Data
	logger                        *log.Helper
	verifierRepo                  *data.VerifierRepo
	transactionRepo               *data.TransactionRepo
	reportTeeAttestationEventRepo *data.ReportTeeAttestationEventRepo
	cAbi                          abi.ABI
	ethClient                     *ethclient.Client
	contractObj                   *contract.Contract

	stopChan             chan struct{}
	blockNumberChan      chan int64
	latestBlockNumber    int64
	processedBlockNumber int64
	processingFlag       bool

	lock sync.RWMutex
}

func NewChain(ctx context.Context, bootstrap *conf.Bootstrap, data *data.Data, logger *log.Helper, verifierRepo *data.VerifierRepo, transactionRepo *data.TransactionRepo, reportTeeAttestationRepo *data.ReportTeeAttestationEventRepo) (*Chain, error) {
	abiFile, err := os.ReadFile(bootstrap.Contract.Abi)
	if err != nil {
		return nil, err
	}
	cAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		return nil, errors.Wrap(err, "abi json error")
	}

	ethClient, err := ethclient.DialContext(ctx, bootstrap.Chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "new eth client error, rpc url: %s", bootstrap.Chain.RpcUrl)
	}

	contractObj, err := contract.NewContract(common.HexToAddress(bootstrap.Contract.Addr), ethClient)
	if err != nil {
		return nil, errors.Wrapf(err, "NewContract error")
	}

	chain := &Chain{
		cf:                            bootstrap,
		data:                          data,
		logger:                        logger,
		verifierRepo:                  verifierRepo,
		transactionRepo:               transactionRepo,
		reportTeeAttestationEventRepo: reportTeeAttestationRepo,
		cAbi:                          cAbi,
		ethClient:                     ethClient,
		contractObj:                   contractObj,

		stopChan:        make(chan struct{}),
		blockNumberChan: make(chan int64, defaultBlockNumberChanLength),
	}

	return chain, nil
}

func (c *Chain) Start(ctx context.Context) error {
	startBlockNumber := c.cf.Chain.StartBlock
	if startBlockNumber == 0 {
		blockNumber, err := c.ethClient.BlockNumber(ctx)
		if err != nil {
			return errors.Wrapf(err, "chain [%s] get BlockNumber error", c.cf.Chain.ChainName)
		}
		startBlockNumber = int64(blockNumber)

		c.lock.Lock()
		c.latestBlockNumber = int64(blockNumber)
		c.lock.Unlock()
	}

	c.logger.WithContext(ctx).Infof("chain [%s] startBlockNumber: %d", c.cf.Chain.ChainName, startBlockNumber)

	c.lock.Lock()
	c.processedBlockNumber = startBlockNumber - 1
	c.blockNumberChan <- startBlockNumber
	c.lock.Unlock()

	go c.run(ctx)
	go c.stick(ctx)

	return nil
}

func (c *Chain) Stop(ctx context.Context) error {
	c.logger.WithContext(ctx).Infof("chain [%s] stopping", c.cf.Chain.ChainName)

	close(c.stopChan)

	// todo stop

	c.logger.WithContext(ctx).Infof("chain [%s] stopped", c.cf.Chain.ChainName)

	return nil
}

func (c *Chain) run(ctx context.Context) {
	for {
		select {
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain run [%s] stopping", c.cf.Chain.ChainName)
			return
		case blockNumber := <-c.blockNumberChan:
			c.Process(ctx, blockNumber)
		}
	}
}

func (c *Chain) stick(ctx context.Context) {
	latestBlockNumberTicker := time.NewTicker(1 * time.Second)
	sendBlockNumberChanTicker := time.NewTicker(1 * time.Second)
	recordStatusTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-c.stopChan:
			c.logger.WithContext(ctx).Infof("chain stick [%s] stopping", c.cf.Chain.ChainName)
			return
		case <-latestBlockNumberTicker.C:
			go c.UpdateLatestBlockNumber(ctx)
		case <-sendBlockNumberChanTicker.C:
			go c.SendBlockNumberChanByTicker(ctx)
		case <-recordStatusTicker.C:
			go c.RecordStatus(ctx)
		}
	}
}

func (c *Chain) RecordStatus(ctx context.Context) {
	c.logger.WithContext(ctx).Infof("chain [%s] status, latestBlockNumber: %d, processedBlockNumber: %d, processingFlag: %v, blockNumberChanLength: %d",
		c.cf.Chain.ChainName, c.GetLatestBlockNumber(), c.GetProcessedBlockNumber(), c.GetProcessingFlag(), c.GetBlockNumberChanLength())
}

func (c *Chain) UpdateLatestBlockNumber(ctx context.Context) {
	blockNumber, err := c.ethClient.BlockNumber(ctx)
	if err != nil {
		c.logger.WithContext(ctx).Errorf("chain [%s] get BlockNumber error: %s", c.cf.Chain.ChainName, err)
	}

	c.lock.Lock()
	c.latestBlockNumber = int64(blockNumber)
	c.lock.Unlock()
}

func (c *Chain) GetLatestBlockNumber() int64 {
	c.lock.RLock()
	blockNumber := c.latestBlockNumber
	c.lock.RUnlock()

	return blockNumber
}

func (c *Chain) SetProcessingFlag(b bool) {
	c.lock.Lock()
	c.processingFlag = b
	c.lock.Unlock()
}

func (c *Chain) GetProcessingFlag() bool {
	c.lock.RLock()
	b := c.processingFlag
	c.lock.RUnlock()

	return b
}

func (c *Chain) SetProcessedBlockNumber(processedBlockNumber int64) {
	c.lock.Lock()
	c.processedBlockNumber = processedBlockNumber
	c.lock.Unlock()
}

func (c *Chain) GetProcessedBlockNumber() int64 {
	c.lock.RLock()
	processedBlockNumber := c.processedBlockNumber
	c.lock.RUnlock()

	return processedBlockNumber
}

func (c *Chain) GetBlockNumberChanLength() int {
	c.lock.RLock()
	blockNumberChanLength := len(c.blockNumberChan)
	c.lock.RUnlock()

	return blockNumberChanLength
}

func (c *Chain) IsBlockNumberChanFull() bool {
	c.lock.RLock()
	blockNumberChanLength := len(c.blockNumberChan)
	c.lock.RUnlock()

	return blockNumberChanLength >= defaultBlockNumberChanLength
}

func (c *Chain) SendBlockNumberChanByTicker(ctx context.Context) {
	if c.GetProcessingFlag() {
		return
	}
	if c.IsBlockNumberChanFull() {
		return
	}

	latestBlockNumber := c.GetLatestBlockNumber()
	processedBlockNumber := c.GetProcessedBlockNumber()

	if latestBlockNumber > processedBlockNumber {
		c.lock.Lock()
		c.blockNumberChan <- processedBlockNumber + 1
		c.lock.Unlock()
	}
}

func (c *Chain) getEndBlockNumber(startBlockNumber int64) int64 {
	endBlockNumber := startBlockNumber

	latestBlockNumber := c.GetLatestBlockNumber()
	if startBlockNumber >= latestBlockNumber {
		return endBlockNumber
	}

	if latestBlockNumber-startBlockNumber > c.cf.Chain.OffsetBlock {
		endBlockNumber = startBlockNumber + c.cf.Chain.OffsetBlock
	} else {
		endBlockNumber = latestBlockNumber - 1
	}

	return endBlockNumber
}

func (c *Chain) SendAttestationTrx(ctx context.Context, attestationIds [][32]byte, results []bool) (string, error) {

	privateKeyBytes, err := Sm4Decrypt(c)
	txHash := ""
	if err != nil {
		return txHash, errors.Wrap(err, "pk decrypt error")
	}
	privateKey, err := crypto.HexToECDSA(string(privateKeyBytes))
	if err != nil {
		return txHash, errors.Wrap(err, "pk sm4 HexToECDSA error")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return txHash, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.ethClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return txHash, errors.Wrap(err, "eth client get Nonce error")
	}
	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return txHash, errors.Wrap(err, "eth client get SuggestGasPrice error")
	}

	chainID, err := c.ethClient.ChainID(ctx)
	if err != nil {
		return txHash, errors.Wrap(err, "eth client get ChainID error")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return txHash, errors.Wrap(err, "bind NewKeyedTransactorWithChainID error")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := c.contractObj.VerifyAttestationBatch(auth, attestationIds, results)
	if err != nil {
		return txHash, errors.Wrap(err, "contract VerifyAttestationBatch error")
	}
	txHash = tx.Hash().Hex()

	c.logger.WithContext(ctx).Infof("tx hash: %s", tx.Hash().Hex())

	return txHash, nil
}

func (c *Chain) process(ctx context.Context, startBlockNumber, endBlockNumber int64) error {
	q := ethereum.FilterQuery{
		FromBlock: big.NewInt(startBlockNumber),
		ToBlock:   big.NewInt(endBlockNumber),
		Addresses: []common.Address{
			common.HexToAddress(c.cf.Contract.Addr),
		},
		Topics: [][]common.Hash{
			{
				common.HexToHash(c.cf.Contract.Topic),
			},
		},
	}

	cLogs, err := c.ethClient.FilterLogs(ctx, q)
	if err != nil {
		return errors.Wrap(err, "client FilterLogs error")
	}

	var logInfoList []LogInfo
	var attestationIds [][32]byte
	var result []bool

	for _, cLog := range cLogs {
		unpackedData, err := c.contractObj.ParseReportTeeAttestation(cLog)
		if err != nil {
			return errors.Wrap(err, "contract ParseReportTeeAttestation error")
		}

		logInfoList = append(logInfoList, LogInfo{
			BlockNumber:        cLog.BlockNumber,
			ContractAddress:    cLog.Address,
			TxHash:             cLog.TxHash,
			TxIndex:            unpackedData.Raw.TxIndex,
			TeeAddress:         unpackedData.TeeAddress,
			CampaignId:         unpackedData.CampaignId,
			AttestationIdBytes: unpackedData.AttestationId,
			AttestationIdStr:   hex.EncodeToString(unpackedData.AttestationId[:]),
			Attestation:        unpackedData.Attestation,
		})

		attestationIds = append(attestationIds, unpackedData.AttestationId)
		// Verify attestation
		isTrue, err := verifyAttestation(c, unpackedData.Attestation)
		if !isTrue {
			result = append(result, false)
			continue
		}
		result = append(result, true)

	}

	c.logger.WithContext(ctx).Infof("logInfoList: %+v", logInfoList)

	// send transaction to chain
	if len(attestationIds) == 0 {
		return nil
	}
	_, err = c.SendAttestationTrx(ctx, attestationIds, result)

	if err != nil {
		return err
	}

	return nil
}

func (c *Chain) Process(ctx context.Context, startBlockNumber int64) {
	processedBlockNumber := c.GetProcessedBlockNumber()
	if processedBlockNumber != startBlockNumber-1 {
		c.logger.WithContext(ctx).Warnf("chain [%s] Process, startBlockNumber: %d, processedBlockNumber: %d", c.cf.Chain.ChainName, startBlockNumber, processedBlockNumber)
		return
	}

	c.SetProcessingFlag(true)
	defer c.SetProcessingFlag(false)

	endBlockNumber := c.getEndBlockNumber(startBlockNumber)

	c.logger.WithContext(ctx).Infof("chain [%s] Process, latestBlockNumber: %d, startBlockNumber: %d, endBlockNumber: %d", c.cf.Chain.ChainName, c.GetLatestBlockNumber(), startBlockNumber, endBlockNumber)

	err := c.process(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		c.logger.WithContext(ctx).Errorf("chain [%s] Process, process error: %s", c.cf.Chain.ChainName, err)
		return
	}

	c.SetProcessedBlockNumber(endBlockNumber)

	if c.IsBlockNumberChanFull() {
		return
	}

	latestBlockNumber := c.GetLatestBlockNumber()
	c.lock.Lock()
	if latestBlockNumber > endBlockNumber {
		c.blockNumberChan <- endBlockNumber + 1
	}
	c.lock.Unlock()
}
