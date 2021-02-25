//+build wireinject

package main

import (
	"context"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/config"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/server"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/google/wire"
)

func InitializeAndRun(ctx context.Context, cfg config.FilePath) (*server.Server, func(), error) {

	panic(
		wire.Build(
			config.NewConfig,
			config.NewServerConfig,
			serviceModule,
			handlersModule,
			server.New,
		),
	)
}

var serviceModule = wire.NewSet(
	service.Module,
	wire.Bind(new(service.Service), new(*service.BookClubService)),
)

var handlersModule = wire.NewSet(
	handlers.Module,
	wire.Bind(new(handlers.Handlers), new(*handlers.BookClubHandler)),
)
