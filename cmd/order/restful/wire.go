//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/irent/internal/adapter/order/restful"
	biz2 "github.com/blackhorseya/irent/internal/app/domain/account/biz"
	repo2 "github.com/blackhorseya/irent/internal/app/domain/account/biz/repo"
	"github.com/blackhorseya/irent/internal/app/domain/order/biz"
	"github.com/blackhorseya/irent/internal/app/domain/order/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/config"
	"github.com/blackhorseya/irent/internal/pkg/httpx"
	"github.com/blackhorseya/irent/internal/pkg/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,
	httpx.ProviderClientSet,

	// server
	httpx.ProviderServerSet,

	// implementation
	restful.OrderSet,
	biz.OrderSet,
	repo.OrderSet,
	biz2.AccountSet,
	repo2.AccountSet,

	// main
	NewService,
)

func CreateService(path string) (*Service, error) {
	panic(wire.Build(providerSet))
}
