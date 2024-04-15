package service

import (
	"context"
	"fmt"
	"github.com/carv-protocol/verifier/api/verifier/v1"
	"github.com/carv-protocol/verifier/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type TransactionService struct {
	v1.UnimplementedTransactionServer

	transaction *biz.TransactionUsecase
	log         *log.Helper
}

func NewTransactionService(trx *biz.TransactionUsecase, logger log.Logger) *TransactionService {
	return &TransactionService{
		transaction: trx,
		log:         log.NewHelper(logger),
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *v1.CreateTransactionRequest) (*v1.CreateTransactionReply, error) {
	transaction, err := s.transaction.CreateTransaction(ctx, req.Hash)
	fmt.Println("err", transaction, err)

	if err != nil {
		return nil, err
	}
	return &v1.CreateTransactionReply{
		Success: true,
	}, nil
}
func (s *TransactionService) UpdateTransaction(ctx context.Context, req *v1.UpdateTransactionRequest) (*v1.UpdateTransactionReply, error) {
	return &v1.UpdateTransactionReply{}, nil
}
func (s *TransactionService) DeleteTransaction(ctx context.Context, req *v1.DeleteTransactionRequest) (*v1.DeleteTransactionReply, error) {
	return &v1.DeleteTransactionReply{}, nil
}
func (s *TransactionService) GetTransaction(ctx context.Context, req *v1.GetTransactionRequest) (*v1.GetTransactionReply, error) {
	return &v1.GetTransactionReply{}, nil
}
func (s *TransactionService) ListTransaction(ctx context.Context, req *v1.ListTransactionRequest) (*v1.ListTransactionReply, error) {
	return &v1.ListTransactionReply{}, nil
}
