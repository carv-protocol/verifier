package worker

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type nodeInf struct {
	nodeId         uint32
	nodeListIndex  uint32
	rewardClaimer  common.Address
	commissionRate uint32
}
type verifyResult struct {
	BlockNumber     uint64
	ContractAddress common.Address
	TxHash          common.Hash
	TxIndex         uint
	requestId       *big.Int
	attestationID   [32]byte
	result          bool
	isReported      bool
}

type TeeReportAttestationEvent struct {
	BlockNumber       uint64
	ContractAddress   common.Address
	TxHash            common.Hash
	TxIndex           uint
	AttestationIDs    [][32]byte
	VerificationInfos []string
	RequestId         *big.Int
}

type VerificationInfo struct {
	AttestationResult uint64 // 0 valid 1 invalid 2 malicious
	Index             uint64
	V                 uint8
	R                 [32]byte
	S                 [32]byte
}

type ConfirmVrfNodesEvent struct {
	BlockNumber     uint64
	ContractAddress common.Address
	TxHash          common.Hash
	TxIndex         uint
	RequestId       *big.Int
	VrfChosen       []uint32
	Deadline        *big.Int
}
