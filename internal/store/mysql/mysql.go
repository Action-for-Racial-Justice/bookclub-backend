package mysql

import (
	"fmt"
	"log"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib" //driver
	"github.com/jmoiron/sqlx"
)

var (
	Module = wire.NewSet(
		New,
	)
)

type (
	Config struct {
		Host     string
		Port     int
		Database string
		User     string
		Password string
	}

	DB struct {
		DB     *sqlx.DB
		config *Config
	}
)

func New(cfg *Config) *DB {

	mysql := &DB{
		config: cfg,
	}

	return mysql
}

//Connect establishes connection to mysql server
func (mysql *DB) Connect() (*sqlx.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		mysql.config.Host, mysql.config.Port, mysql.config.User, mysql.config.Password, mysql.config.Database)
	log.Println("Connection string -->", connString)
	mysqlDB, err := sqlx.Open("mysql", connString)
	if err != nil && mysqlDB != nil {
		log.Println("Error connecting to database")
		log.Println(err)
		return nil, err
	}
	err = mysqlDB.Ping()
	if err != nil {
		log.Println(fmt.Sprintf("mysql ping failed on startup, will keep trying. Error was %+v", err))
	}
	return mysqlDB, nil
}

//Close closes current connection w/ mysql server
func (postgres *DB) Close() error {
	if postgres != nil {
		err := postgres.DB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
