//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/carv-protocol/verifier/internal/biz"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/data"
	"github.com/carv-protocol/verifier/internal/server"
	"github.com/carv-protocol/verifier/internal/service"
	"github.com/carv-protocol/verifier/internal/worker"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, *log.Helper) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, worker.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
