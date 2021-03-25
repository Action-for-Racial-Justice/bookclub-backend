package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_USER_DATA_QUERY = `SELECT`

func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	result, err := sql.db.db.Query("SELECT * from user u WHERE u.id='%s'", userID)
	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	var userData models.UserData
	if err = result.Scan(&userData.ID, &userData.FullName, &userData.ClubAssigned); err != nil {
		log.Printf("error while scanning result for user data: %s", err)
		return nil, err
	}

	return &userData, nil

}
