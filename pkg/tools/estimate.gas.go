package tools

import (
	"github.com/carv-protocol/verifier/pkg/contract"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func EstimateGas() {
	abi := contract.ContractMetaData.ABI

	// Node Enter Signature
	NODE_ENTER_SIGNATURE := []byte("nodeEnter(address)")
	hash := sha3.N
	var data []byte

	data = append(data, NODE_ENTER_SIGNATURE...)
}
