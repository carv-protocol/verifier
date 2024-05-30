package worker

import "github.com/ethereum/go-ethereum/common"

type LogInfo struct {
	BlockNumber        uint64
	ContractAddress    common.Address
	TxHash             common.Hash
	TxIndex            uint
	TeeAddress         common.Address
	CampaignId         string
	AttestationIdBytes [32]byte
	AttestationIdStr   string
	Attestation        string
}
