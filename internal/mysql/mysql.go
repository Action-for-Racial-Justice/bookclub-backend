//go:generate mockgen -package=mocks -destination=../mocks/mysql.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql Mysql

package mysql

import (
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	_ "github.com/go-sql-driver/mysql" //driver
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

//Module to denote wire binding function
var (
	Module = wire.NewSet(
		New,
	)
)

type (
	//Mysql interface which describes BookClubMysql struct functions (currently none)
	Mysql interface {
		CreateUserClubMember(joinRequest *models.JoinClubRequest) error
		DeleteClub(deleteRequest *models.LeaveClubRequest) error
		DeleteUserClubMember(deleteRequest *models.LeaveClubRequest) error
		GetUserDataForUserID(userID string) (*models.UserData, error)
		GetUserClubMembers(userID string) ([]models.ClubMember, error)
		GetUserClubs([]models.ClubMember) (*models.Clubs, error)
		GetClubDataForEntryID(entryID string) (*models.Club, error)
		GetBookDataForEntryID(entryID string) (*models.Book, error)
		GetListClubs() (*models.Clubs, error)
		IsUserClubLeader(*models.LeaveClubRequest) (bool, error)
		UserJoinClub(*models.JoinClubRequest) (string, error)
		InsertUser(*models.ArjUser) error
		CreateClub(createRequest *models.CreateClubRequest) error
	}

	//BookClubMysql struct to hold relevant inner data members and functions for database connection
	BookClubMysql struct {
		mysql *DB
	}

	//Config ... Configuration struct
	Config struct {
		Host     string
		Port     int
		Database string
		User     string
		Password string
	}

	//DB inner data struct for our database, holds configuration and functions for db connection
	DB struct {
		db     *sqlx.DB
		config *Config
	}
)

//New ... constructor
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
	bcMysql := &BookClubMysql{
		mysql: database,
	}

	log.Println("Connecting to Database")
	dbConnection, err := database.Connect()
	if err != nil {
		return nil, nil, err
	}

	bcMysql.mysql.db = dbConnection

	return bcMysql, close, nil
}

//Connect opens a database specified by its database driver name, and a connection string which contains relevant connection information
func (mysql *DB) Connect() (*sqlx.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysql.config.User, mysql.config.Password, mysql.config.Host, mysql.config.Port, mysql.config.Database)

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

//Close closes a database and prevents new queries from starting
func (mysql *DB) Close() error {
	if mysql != nil {
		err := mysql.db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//Closes Statement
func closeStatement(stmt *sqlx.Stmt) {
	if err := stmt.Close(); err != nil {
		log.Printf("Error trying to close statement: %s", err.Error())
	}
}

//Closes Named Statement
func closeNamedStatement(stmt *sqlx.NamedStmt) {
	if err := stmt.Close(); err != nil {
		log.Printf("Error trying to close statement: %s", err.Error())
	}
}

//Closes Rows
func closeRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Printf("Error trying to close statement: %s", err.Error())

	}
}
