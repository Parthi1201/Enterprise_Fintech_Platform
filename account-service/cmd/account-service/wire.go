//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"account-service/internal/handler"
	"account-service/internal/conf"
	"account-service/internal/data"
	"account-service/internal/server"
	"account-service/internal/service"
	"account-service/internal/client"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, handler.HandlerProviderSet, service.ServiceProviderSet,client.ClientProviderSet ,newApp))
}
