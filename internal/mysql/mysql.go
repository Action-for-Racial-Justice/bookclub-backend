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
	MySqlConnection interface {
		MySQLConnect() error
	}

	Database struct { //conduitstroage
		mysqlDB *DB //postgresconnections
	}

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

func New(cfg *Config) (*Database, func()) {

	tempDB := &DB{
		config: cfg,
	}

	close := func() {
		log.Println("shutting down mysql connection")
		if tempDB.DB == nil {
			return
		}

		if err := tempDB.Close(); err != nil {
			log.Printf("Error occured closing DB connection: %+v", err)
		}
		log.Println("mysql connection shutdown")
	}
	datab := &Database{
		mysqlDB: tempDB,
	}

	log.Println("Connecting to Database")
	db, err := tempDB.Connect()
	if err != nil {
		return nil, nil
	}

	datab.mysqlDB.DB = db

	return datab, close
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
func (mysql *DB) Close() error {
	if mysql != nil {
		err := mysql.DB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
