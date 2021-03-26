package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_USER_DATA_QUERY = "SELECT * FROM user where id = ?"

func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	stmt, err := sql.db.db.Preparex(GET_USER_DATA_QUERY)
	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(1) //can't pass in string, can't even convert to int and then pass in, only works like this

	var userData models.UserData
	if err = row.StructScan(&userData); err != nil {
		log.Printf("error while scanning result for user data: %s", err)
		return nil, err
	}

	return &userData, nil

}
