package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_USER_DATA_QUERY = "SELECT * FROM user where id = ?"
const CREATE_USER_CLUB_MEMBER = "INSERT INTO club_member(id, uid, clubId) VALUES(:ID, :userID, :clubID)"

func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.UserData, error) {

	stmt, err := sql.db.db.Preparex(GET_USER_DATA_QUERY)
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

func (sql *BookClubMysql) CreateUserClubMember(clubMember *models.JoinClubRequest) error {
	stmt, err := sql.db.db.PrepareNamed(CREATE_USER_CLUB_MEMBER)
	if err != nil {
		log.Printf("error while preparing user clum member insert: %s", err)
		return err
	}

	result, err := stmt.Exec(clubMember)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows == 0 && err == nil {
		err = errors.New(fmt.Sprintf("user already exist in club %s", clubMember.ClubID))
		log.Print(err.Error())
		return err
	}

	return nil
}
