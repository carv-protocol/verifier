package tools

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func EstimateGas(rpcUrl, from, to, data string, value uint64) uint64 {

	// Establish an RPC connection to the specified RPC url
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalln(err)
	}

	var ctx = context.Background()

	var (
		fromAddr  = common.HexToAddress(from)     // Convert the from address from hex to an Ethereum address.
		toAddr    = common.HexToAddress(to)       // Convert the to address from hex to an Ethereum address.
		amount    = new(big.Int).SetUint64(value) // Convert the value from uint64 to *big.Int.
		bytesData []byte
	)

	// Encode the data if it's not already hex-encoded.
	if data != "" {
		if ok := strings.HasPrefix(data, "0x"); !ok {
			data = hexutil.Encode([]byte(data))
		}

		bytesData, err = hexutil.Decode(data)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Create a message which contains information about the transaction.
	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	// Estimate the gas required for the transaction.
	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}

	return gas
}
