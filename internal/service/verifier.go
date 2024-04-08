package service

import (
	"context"

	v1 "github.com/carv-protocol/verifier/api/verifier/v1"
	"github.com/carv-protocol/verifier/internal/biz"
)

type VerifierService struct {
	v1.UnimplementedVerifierServer

	uc *biz.VerifierUsecase
}

func NewVerifierService(uc *biz.VerifierUsecase) *VerifierService {
	return &VerifierService{uc: uc}
}

func (s *VerifierService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateVerifier(ctx, &biz.Verifier{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
