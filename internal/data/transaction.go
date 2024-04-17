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

func (t *TransactionRepo) CreateTransaction(ctx context.Context, transaction *biz.Transaction) (*biz.Transaction, error) {
	t.data.db.Create(&transaction)
	t.log.WithContext(ctx).Info("gormDB: CreateTransaction, trx:", transaction.TxHash)
	return transaction, nil
}
