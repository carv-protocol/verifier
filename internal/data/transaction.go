package data

import (
	"context"
	"github.com/carv-protocol/verifier/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type TransactionRepo struct {
	data *Data
	log  *log.Helper
}

func NewTransactionRepo(data *Data, logger log.Logger) *TransactionRepo {
	return &TransactionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *TransactionRepo) CreateTransaction(ctx context.Context, hash string, createdTime int64, updatedTime int64) (*biz.Transaction, error) {
	var stu biz.Transaction
	stu.Hash = hash
	stu.CreatedAt = createdTime
	stu.UpdatedAt = updatedTime
	t.data.db.Create(&stu)
	t.log.WithContext(ctx).Info("gormDB: CreateTransaction, trx:", hash)
	return &stu, nil
}
