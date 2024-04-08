package data

import (
	"context"

	"github.com/carv-protocol/verifier/internal/infra/data/dal"
	"github.com/carv-protocol/verifier/internal/infra/data/model"
)

type VerifierRepo struct {
	data *Data
}

func NewVerifierRepo(data *Data) *VerifierRepo {
	return &VerifierRepo{
		data: data,
	}
}

func (r *VerifierRepo) ListAll(ctx context.Context) ([]*model.Verifier, error) {
	q := dal.Use(r.data.DB(ctx)).Verifier

	return q.WithContext(ctx).Where().Find()
}
