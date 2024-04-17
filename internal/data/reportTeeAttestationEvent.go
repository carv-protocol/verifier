package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/carv-protocol/verifier/internal/biz"
)

type ReportTeeAttestationEventRepo struct {
	data *Data
	log  *log.Helper
}

func NewReportTeeAttestationEventRepo(data *Data, logger log.Logger) *ReportTeeAttestationEventRepo {
	return &ReportTeeAttestationEventRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *ReportTeeAttestationEventRepo) CreateReportTeeAttestationEvent(ctx context.Context, event biz.ReportTeeAttestationEvent) (*biz.ReportTeeAttestationEvent, error) {
	t.log.WithContext(ctx).Info("gormDB: CreateReportTeeAttestationEvent")
	t.data.db.Create(&event)
	return &event, nil
}

func (t *ReportTeeAttestationEventRepo) FindReportTeeAttestationEvent(ctx context.Context, hash string, txIndex int) (*biz.ReportTeeAttestationEvent, *gorm.DB) {
	var event *biz.ReportTeeAttestationEvent
	result := t.data.db.Find(&event, "tx_hash = ? AND tx_index = ?", hash, txIndex)
	return event, result
}

func (t *ReportTeeAttestationEventRepo) LastReportTeeAttestationEvent(rtae *biz.ReportTeeAttestationEvent) (*biz.ReportTeeAttestationEvent, error) {

	err := t.data.db.Last(&rtae).Error
	if err != nil {
		return nil, err
	}
	return rtae, nil
}
