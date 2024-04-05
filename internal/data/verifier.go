package data

import (
	"context"
)

type VerifierRepo struct {
	data *Data
}

func NewVerifierRepo(data *Data) *VerifierRepo {
	return &VerifierRepo{
		data: data,
	}
}

func (r *VerifierRepo) ListAll(context.Context) (interface{}, error) {
	return nil, nil
}
