//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"slate/internal/controller"
	"slate/internal/databasePool"
	"slate/internal/repo"
	"slate/internal/router"
	"slate/internal/server"
	"slate/internal/service"

	"github.com/google/wire"
)

func InitializeApp(ctx context.Context, connString string) (*server.Server, error) {
	wire.Build(router.ProviderSetRouter, server.ProviderSetHTTPServer, controller.ProviderSetController, service.ProviderSetService, repo.ProviderSetRepo, databasePool.ProviderSetDatabasePool)
	return &server.Server{}, nil
}
