package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/carv-protocol/verifier/api/verifier/v1"
	"github.com/carv-protocol/verifier/internal/infra/data/model"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Verifier is a Verifier model.
type Verifier struct {
	Hello string
}

type VerifierRepo interface {
	ListAll(ctx context.Context) ([]*model.Verifier, error)
}

// VerifierUsecase is a Verifier usecase.
type VerifierUsecase struct {
	repo   VerifierRepo
	logger *log.Helper
}

// NewVerifierUsecase new a Verifier usecase.
func NewVerifierUsecase(repo VerifierRepo, logger *log.Helper) *VerifierUsecase {
	return &VerifierUsecase{repo: repo, logger: logger}
}

// CreateVerifier creates a Verifier, and returns the new Verifier.
func (uc *VerifierUsecase) CreateVerifier(ctx context.Context, g *Verifier) (*Verifier, error) {
	uc.logger.WithContext(ctx).Infof("CreateVerifier: %v", g.Hello)
	return g, nil
}
