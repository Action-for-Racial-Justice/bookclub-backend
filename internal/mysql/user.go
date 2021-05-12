package mysql

import (
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	createClubMemberQuery = "INSERT IGNORE INTO club_member(entryID, userID, clubID) VALUES(:entryID, :userID, :clubID);"
	deleteClubMemberQuery = "DELETE FROM club_member WHERE userID = :userID AND clubID = :clubID;"
	getUserDataQuery      = "SELECT * FROM user where id = ?;"
	isUserLeaderQuery     = "SELECT COUNT(*) FROM club WHERE leaderID = :userID AND entryID = :clubID;"
	insertUserQuery       = "INSERT INTO user(id, fullName) VALUES(:id, :fullName);"
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

//IsUserClubLeader returns true if the user is the leader of the club
func (bcm *BookClubMysql) IsUserClubLeader(leaveRequest *models.LeaveClubRequest) (bool, error) {

	var count int

	stmt, err := bcm.mysql.db.PrepareNamed(isUserLeaderQuery)
	defer closeNamedStatement(stmt)

	if err != nil {
		log.Printf("error while checking if user is leader of club: %s", err)
		return false, err
	}

	result := stmt.QueryRowx(leaveRequest)

	err = result.Scan(&count)

	if err != nil {
		log.Printf("error while executing isUserLeaderQuery: %s", err)
		return false, err
	} else if count == 0 {
		log.Println("user is not leader")
		return false, nil
	} else {
		log.Println("user is leader, returning true")
		return true, nil
	}
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

//DeleteUserClubMember deletes a single club member entry for the user
func (bcm *BookClubMysql) DeleteUserClubMember(clubMember *models.LeaveClubRequest) error {
	stmt, err := bcm.mysql.db.PrepareNamed(deleteClubMemberQuery)

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
				"could not delete club member entry %s",
				clubMember.ClubID,
			),
			bcerrors.InternalError,
		)
	}
	return nil
}

//InsertUser inserts user to database
func (bcm *BookClubMysql) InsertUser(user *models.ArjUser) error {
	stmt, err := bcm.mysql.db.PrepareNamed(insertUserQuery)

	if err != nil {
		log.Printf("error while preparing user insert: %s", err)
		return err
	}
	defer closeNamedStatement(stmt)

	if _, err = stmt.Exec(user); err != nil {
		log.Printf("error while executing user insert: %s", err)
		return err
	}
	return nil
}
