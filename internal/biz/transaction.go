package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Transaction struct {
	ID        int32
	Hash      string
	Status    int32
	CreatedAt int64
	UpdatedAt int64
}

type TransactionRepo interface {
	CreateTransaction(context.Context, string, int64, int64) (*Transaction, error)
}

type TransactionUsecase struct {
	repo TransactionRepo
	log  *log.Helper
}

func NewTransactionUsecase(repo TransactionRepo, logger *log.Helper) *TransactionUsecase {
	return &TransactionUsecase{repo: repo, log: logger}
}

func (trx *TransactionUsecase) CreateTransaction(ctx context.Context, hash string) (*Transaction, error) {
	trx.log.WithContext(ctx).Infof("CreateTransaction: %v", hash)
	return trx.repo.CreateTransaction(ctx, hash, time.Now().Unix(), time.Now().Unix())
}
