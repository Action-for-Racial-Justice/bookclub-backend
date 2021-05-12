// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/config"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/requests"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/server"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/validator"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitializeAndRun(ctx context.Context, cfg config.FilePath) (*server.Server, func(), error) {
	configConfig := config.NewConfig(cfg)
	serverConfig := config.NewServerConfig(configConfig)
	mysqlConfig := config.NewDBConfig(configConfig)
	bookClubMysql, cleanup, err := mysql.New(mysqlConfig)
	if err != nil {
		return nil, nil, err
	}
	requestsConfig := config.NewRequestsConfig(configConfig)
	requestsRequests := requests.New(requestsConfig)
	bcValidator := validator.New()
	v := config.NewLoggerOptions()
	logger, err := zap.NewProduction(v...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	bookClubService := service.New(bookClubMysql, requestsRequests, bcValidator, logger)
	bookClubHandler, err := handlers.New(bookClubService, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serverServer, cleanup2, err := server.New(ctx, serverConfig, bookClubHandler)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return serverServer, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var databaseModule = wire.NewSet(mysql.Module, wire.Bind(new(mysql.Mysql), new(*mysql.BookClubMysql)))

var handlersModule = wire.NewSet(handlers.Module, wire.Bind(new(handlers.Handlers), new(*handlers.BookClubHandler)))

var requestsModule = wire.NewSet(requests.Module, wire.Bind(new(requests.IRequests), new(*requests.Requests)))

var serviceModule = wire.NewSet(service.Module, wire.Bind(new(service.Service), new(*service.BookClubService)))

var validatorModule = wire.NewSet(validator.Module, wire.Bind(new(validator.Validator), new(*validator.BCValidator)))
