package worker

import (
	"context"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewWorkerServer)

type Server struct {
	cf           *conf.Bootstrap
	data         *data.Data
	logger       *log.Helper
	verifierRepo *data.VerifierRepo
	chain        *Chain
}

func NewWorkerServer(bootstrap *conf.Bootstrap, data *data.Data, logger *log.Helper, verifierRepo *data.VerifierRepo) *Server {
	w := &Server{
		cf:           bootstrap,
		data:         data,
		logger:       logger,
		verifierRepo: verifierRepo,
	}

	return w
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.WithContext(ctx).Info("worker server starting")

	chain, err := NewChain(ctx, s.cf, s.data, s.logger, s.verifierRepo)
	if err != nil {
		return err
	}

	if err = chain.Start(ctx); err != nil {
		return err
	}

	s.chain = chain

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.WithContext(ctx).Info("worker server stopping")

	if err := s.chain.Stop(ctx); err != nil {
		s.logger.WithContext(ctx).Error(err)
	}

	s.logger.WithContext(ctx).Info("worker server stopped")
	return nil
}
