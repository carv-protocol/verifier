package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ReportTeeAttestationEvent struct {
	ID              int32
	BlockNumber     uint64
	ContractAddress string
	TeeAddress      string
	CampaignId      string
	AttestationId   []byte
	Attestation     string
	TxHash          string
	TxIndex         int
	CreatedAt       int32
	HandleAt        int32
}

type ReportTeeAttestationEventRepo interface {
	CreateReportTeeAttestationEvent(context.Context, ReportTeeAttestationEvent) (*ReportTeeAttestationEvent, error)
}

type ReportTeeAttestationEventUsecase struct {
	repo ReportTeeAttestationEventRepo
	log  *log.Helper
}

func NewReportTeeAttestationEventUsecase(repo ReportTeeAttestationEventRepo, logger *log.Helper) *ReportTeeAttestationEventUsecase {
	return &ReportTeeAttestationEventUsecase{repo: repo, log: logger}
}

func (rtaeUsecase *ReportTeeAttestationEventUsecase) CreateReportTeeAttestationEvent(ctx context.Context, event ReportTeeAttestationEvent) (*ReportTeeAttestationEvent, error) {
	rtaeUsecase.log.WithContext(ctx).Infof("CreateTransaction: %v", event)

	return rtaeUsecase.repo.CreateReportTeeAttestationEvent(ctx, event)

}
