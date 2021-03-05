package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //driver
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var (
	Module = wire.NewSet(
		New,
	)
)

type (
	Mysql interface {
	}

	BookClubMysql struct {
		mysqlDB *DB
	}

	Config struct {
		Host     string
		Port     int
		Database string
		User     string
		Password string
	}

	DB struct {
		db     *sqlx.DB
		config *Config
	}
)

func New(cfg *Config) (*BookClubMysql, func(), error) {

	database := &DB{
		config: cfg,
	}

	close := func() {
		log.Println("shutting down mysql connection")
		if database.db == nil {
			return
		}

		if err := database.Close(); err != nil {
			log.Printf("Error occured closing DB connection: %+v", err)
		}
		log.Println("mysql connection shutdown")
	}
	mysql := &BookClubMysql{
		mysqlDB: database,
	}

	log.Println("Connecting to Database")
	dbConnection, err := database.Connect()
	if err != nil {
		return nil, nil, err
	}

	mysql.mysqlDB.db = dbConnection

	return mysql, close, nil
}

func (mysql *DB) Connect() (*sqlx.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysql.config.User, mysql.config.Password, mysql.config.Host, mysql.config.Port, mysql.config.Database)

	log.Println("Connection string -->", connString)

	mysqlDB, err := sqlx.Open("mysql", connString)
	if err != nil && mysqlDB != nil {
		log.Println("Error connecting to database")
		log.Println(err)
		return nil, err
	}

	if err = mysqlDB.Ping(); err != nil {
		log.Println(fmt.Sprintf("mysql ping failed on startup, will keep trying. Error was %+v", err))
	}

	return mysqlDB, nil
}

func (mysql *DB) Close() error {
	if mysql != nil {
		err := mysql.db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
