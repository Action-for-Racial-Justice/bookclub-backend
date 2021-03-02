package config

import (
	"log"
	"strconv"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/server"
	"github.com/joho/godotenv"

	"os"
)

type FilePath string

type Config struct {
	serverConfig server.Config
	DbConfig     mysql.Config
}

func NewConfig(fileName FilePath) *Config {

	if err := godotenv.Load(string(fileName)); err != nil {
		log.Printf("Config file not found for file name: %s", fileName)
		panic(err)
	}

	return &Config{
		serverConfig: server.Config{
			Host:            os.Getenv("SERVER_HOST"),
			Port:            convertToInt(os.Getenv("SERVER_PORT")),
			ListenLimit:     convertToInt(os.Getenv("SERVER_LISTEN_LIMIT")),
			KeepAlive:       convertToInt(os.Getenv("SERVER_KEEP_ALIVE_TIME")),
			ReadTimeout:     convertToInt(os.Getenv("SERVER_READ_TIMEOUT")),
			WriteTimeout:    convertToInt(os.Getenv("SERVER_WRITE_TIMEOUT")),
			ShutdownTimeout: convertToInt(os.Getenv("SERVER_SHUTDOWN_TIME")),
		},
		DbConfig: mysql.Config{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     convertToInt(os.Getenv("MYSQL_PORT")),
			Database: os.Getenv("DATABASE"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("PASSWORD"),
		},
	}
}

func NewServerConfig(cfg *Config) *server.Config {
	return &cfg.serverConfig
}

func NewDBConfig(cfg *Config) *mysql.Config {
	return &cfg.DbConfig
}

func convertToInt(str string) int {
	intRep, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return intRep
}
