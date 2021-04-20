package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

//GetUserDataForUserID returns userData struct holding bookclub user data for a userID string
func (bcm *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	stmt, err := bcm.mysql.db.Preparex(GET_USER_DATA_QUERY)
	defer closeStatement(stmt)

	if err != nil {
		log.Printf("error while querying db for user data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(userID)
	var userData models.UserData
	if err = row.StructScan(&userData); err != nil {
		log.Printf("error while scanning result for user data: %s", err)
		return nil, err
	}

	return &userData, nil

}

//CreateUserClubMember creates club member in the clubmember mysql table s
func (bcm *BookClubMysql) CreateUserClubMember(clubMember *models.JoinClubRequest) error {
	stmt, err := bcm.mysql.db.PrepareNamed(CREATE_USER_CLUB_MEMBER)

	if err != nil {
		log.Printf("error while preparing user club member insert: %s", err)
		return err
	}

	defer closeNamedStatement(stmt)

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
