package service

import (
	"context"
	v1 "github.com/carv-protocol/verifier/api/verifier/v1"
	"github.com/carv-protocol/verifier/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ReportTeeAttestationEventService struct {
	reportTeeAttestationEvent *biz.ReportTeeAttestationEventUsecase
	log                       *log.Helper
}

func NewReportTeeAttestationEventService(rtae *biz.ReportTeeAttestationEventUsecase, logger log.Logger) *ReportTeeAttestationEventService {
	return &ReportTeeAttestationEventService{
		reportTeeAttestationEvent: rtae,
		log:                       log.NewHelper(logger),
	}

}

func (rtae *ReportTeeAttestationEventService) CreateReportTeeAttestationEvent(ctx context.Context, event biz.ReportTeeAttestationEvent) (*v1.CreateReportTeeAttestationEventReply, error) {
	_, err := rtae.reportTeeAttestationEvent.CreateReportTeeAttestationEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return &v1.CreateReportTeeAttestationEventReply{
		Success: true,
	}, nil
}
