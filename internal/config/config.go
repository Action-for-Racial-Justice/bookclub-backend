package config

import (
	"log"
	"strconv"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/requests"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/server"

	"github.com/joho/godotenv"

	"os"
)

//FilePath struct to be propogated through wire
type FilePath string

//Config ...
type Config struct {
	requestsConfig requests.Config
	serverConfig   server.Config
	dbConfig       mysql.Config
}

//NewConfig builds global config struct
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
		dbConfig: mysql.Config{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     convertToInt(os.Getenv("MYSQL_PORT")),
			Database: os.Getenv("MYSQL_DATABASE"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
		requestsConfig: requests.Config{
			ArjBackendURL: os.Getenv("ARJ_BACKEND_URL"),
		},
	}
}

//NewServerConfig returns server config from global config
func NewServerConfig(cfg *Config) *server.Config {
	return &cfg.serverConfig
}

//NewDBConfig returns database config from global config
func NewDBConfig(cfg *Config) *mysql.Config {
	return &cfg.dbConfig
}

//NewRequestsConfig returns requests config from global config
func NewRequestsConfig(cfg *Config) *requests.Config {
	return &cfg.requestsConfig
}

func convertToInt(str string) int {
	intRep, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return intRep
}
