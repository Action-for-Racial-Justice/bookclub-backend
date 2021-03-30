//+build wireinject

package main

import (
	"context"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/config"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/server"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/validator"

	"github.com/google/wire"
)

func InitializeAndRun(ctx context.Context, cfg config.FilePath) (*server.Server, func(), error) {

	panic(
		wire.Build(
			config.NewConfig,
			config.NewServerConfig,
			config.NewDBConfig,
			databaseModule,
			validatorModule,
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

var databaseModule = wire.NewSet(
	mysql.Module,
	wire.Bind(new(mysql.Mysql), new(*mysql.BookClubMysql)),
)

var validatorModule = wire.NewSet(
	validator.Module,
	wire.Bind(new(validator.Validator), new(*validator.BCValidator)),
)
