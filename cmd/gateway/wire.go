//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gateway/internal/api"
	"gateway/internal/client"
	"gateway/internal/conf"
	"gateway/internal/gin"
	"gateway/internal/pkg/registry"
	"gateway/internal/server"
	"gateway/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Etcd, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(registry.ProviderSet, server.ProviderSet, gin.ProviderSet, client.ProviderSet, service.ProviderSet, api.ProviderSet, newApp))
}
