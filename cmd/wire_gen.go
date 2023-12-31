// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"slate/internal/controller"
	"slate/internal/databasePool"
	"slate/internal/repo"
	"slate/internal/router"
	"slate/internal/server"
	"slate/internal/service"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, connString string) (*server.Server, error) {
	database, err := databasePool.NewDatabase(ctx, connString)
	if err != nil {
		return nil, err
	}
	userRepo := repo.NewUserRepo(database, ctx)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService)
	userController := controller.NewUserController(userService, authService)
	apiRouter := router.NewAPIRouter(userController)
	serverServer := server.NewHTTPServer(ctx, apiRouter)
	return serverServer, nil
}
