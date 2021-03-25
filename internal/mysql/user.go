package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_USER_DATA_QUERY = "SELECT * from user WHERE id=?"

func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	stmt, err := sql.db.db.Preparex(GET_USER_DATA_QUERY)
	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	rows, err := stmt.Exec(userID)

	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	var userData models.UserData
	if err = row.Scan(&userData.ID, &userData.FullName, &userData.ClubAssigned); err != nil {
		log.Printf("error while scanning result for user data: %s", err)
		return nil, err
	}

	return &userData, nil

}
