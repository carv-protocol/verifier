package worker

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/data"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

const (
	defaultBlockNumberChanLength = 10
)

type Chain struct {
	cf        *conf.Bootstrap
	data      *data.Data
	logger    *log.Helper
	ethClient *ethclient.Client
	cAbi      abi.ABI

	stopChan             chan struct{}
	blockNumberChan      chan int64
	latestBlockNumber    int64
	processedBlockNumber int64
	processingFlag       bool

	lock sync.RWMutex
}

func NewChain(ctx context.Context, bootstrap *conf.Bootstrap, data *data.Data, logger *log.Helper) (*Chain, error) {
	cAbi, err := abi.JSON(strings.NewReader(bootstrap.Contract.Abi))
	if err != nil {
		return nil, errors.Wrap(err, "abi json error")
	}

	ethClient, err := ethclient.DialContext(ctx, bootstrap.Chain.RpcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "new eth client error, rpc url: %s", bootstrap.Chain.RpcUrl)
	}

	chain := &Chain{
		cf:              bootstrap,
		data:            data,
		logger:          logger,
		ethClient:       ethClient,
		cAbi:            cAbi,
		stopChan:        make(chan struct{}),
		blockNumberChan: make(chan int64, defaultBlockNumberChanLength),
	}

	return chain, nil
}

func (c *Chain) Start(ctx context.Context) error {
	startBlockNumber := c.cf.Chain.StartBlock
	// todo read from db
	//if record != nil && record.BlockNumber > startBlockNumber {
	//	startBlockNumber = record.BlockNumber + 1
	//}
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
	for _, cLog := range cLogs {
		unpackedData, err := c.cAbi.Unpack("ReportTeeAttestation", cLog.Data)
		if err != nil {
			return errors.Wrap(err, "abi Unpack error")
		}
		campaignId, ok := unpackedData[1].(string)
		if !ok {
			return fmt.Errorf("log data parse campaignId not ok, data: %v", unpackedData)
		}
		attestationIdBytes, ok := unpackedData[2].([32]byte)
		if !ok {
			return fmt.Errorf("log data parse attestationId not ok, data: %v", unpackedData)

		}
		attestationIdStr := hex.EncodeToString(attestationIdBytes[:])
		attestation, ok := unpackedData[3].(string)
		if !ok {
			return fmt.Errorf("log data parse attestation not ok, data: %v", unpackedData)
		}

		logInfoList = append(logInfoList, LogInfo{
			BlockNumber:        cLog.BlockNumber,
			Address:            cLog.Address,
			TxHash:             cLog.TxHash,
			CampaignId:         campaignId,
			AttestationIdBytes: attestationIdBytes,
			AttestationIdStr:   attestationIdStr,
			Attestation:        attestation,
		})
	}

	c.logger.WithContext(ctx).Infof("logInfoList: %+v", logInfoList)

	// todo verify and report

	// todo db store

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
