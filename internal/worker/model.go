package worker

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type nodeInf struct {
	nodeId         uint32
	nodeListIndex  uint32
	rewardClaimer  common.Address
	commissionRate uint32
}
type verifyResult struct {
	attestationID [32]byte
	result        bool
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
