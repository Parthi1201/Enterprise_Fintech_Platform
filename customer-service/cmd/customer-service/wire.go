//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"customer-service/internal/handler"
	"customer-service/internal/conf"
	"customer-service/internal/data"
	"customer-service/internal/server"
	"customer-service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, handler.HandlerProviderSet, service.ServiceProviderSet, newApp))
}
