package worker

import (
	"context"

	"github.com/carv-protocol/verifier/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewWorkerServer)

type Server struct {
	cf     *conf.Bootstrap
	logger *log.Helper
	chain  *Chain
}

func NewWorkerServer(bootstrap *conf.Bootstrap, logger *log.Helper) *Server {
	w := &Server{
		cf:     bootstrap,
		logger: logger,
	}

	return w
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.WithContext(ctx).Info("worker server starting")

	chain, err := NewChain(ctx, s.cf, s.logger)
	if err != nil {
		return err
	}

	for i := 0; i < len(s.cf.Chains); i++ {
		if err = chain.Start(ctx, i); err != nil {
			return err
		}

	}

	s.chain = chain

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.WithContext(ctx).Info("worker server stopping")

	for i := 0; i < len(s.cf.Chains); i++ {
		if err := s.chain.Stop(ctx, i); err != nil {
			s.logger.WithContext(ctx).Error(err)
		}
	}

	s.logger.WithContext(ctx).Info("worker server stopped")
	return nil
}
