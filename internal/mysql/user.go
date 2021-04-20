package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserDataResponse, error) {

	stmt, err := sql.db.db.Preparex(GET_USER_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(userID)
	var userData models.UserDataResponse
	if err = row.StructScan(&userData); err != nil {
		log.Printf("error while scanning result for user data: %s", err)
		return nil, err
	}

	return &userData, nil

}

func (sql *BookClubMysql) CreateUserClubMember(clubMember *models.JoinClubRequest) error {
	stmt, err := sql.db.db.PrepareNamed(CREATE_USER_CLUB_MEMBER)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while preparing user club member insert: %s", err)
		return err
	}

	result, err := stmt.Exec(clubMember)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		err = errors.New(fmt.Sprintf("user already exist in club %s", clubMember.ClubID))
		log.Print(err.Error())
		return err
	}

	return nil
}
