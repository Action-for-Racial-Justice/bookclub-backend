package mysql

import (
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	createClubMemberQuery = "INSERT INTO club_member(entryID, userID, clubID) VALUES(:entryID, :userID, :clubID)"
	getUserDataQuery      = "SELECT * FROM user where id = ?"
)

//GetUserDataForUserID returns userData struct holding bookclub user data for a userID string
func (bcm *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	stmt, err := bcm.mysql.db.Preparex(getUserDataQuery)
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
	stmt, err := bcm.mysql.db.PrepareNamed(createClubMemberQuery)

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
		return bcerrors.NewError(
			fmt.Sprintf(
				"user already exist in club %s",
				clubMember.ClubID,
			),
			bcerrors.InternalError,
		)
	}
	return nil
}
