//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"card-service/internal/handler"
	"card-service/internal/conf"
	"card-service/internal/data"
	"card-service/internal/server"
	"card-service/internal/service"
	"card-service/internal/client"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, handler.HandlerProviderSet, service.ServiceProviderSet, client.ClientProviderSet,newApp))
}
