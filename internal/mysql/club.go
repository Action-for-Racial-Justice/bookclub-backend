package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	CREATE_USER_CLUB_MEMBER          = "INSERT INTO club_member(entryID, userID, clubID) VALUES(:entryID, :userID, :clubID)"
	GET_CLUBS_DATA_QUERY             = "SELECT * FROM club"
	GET_USER_CLUB_MEMBERS_DATA_QUERY = "SELECT * FROM club_member where userID = ?"
	GET_CLUB_DATA_QUERY              = "SELECT * FROM club where entryID = ?"
	CREATE_CLUB                      = "INSERT INTO club(entryID, leaderID, clubName, bookID) VALUES(:entryID, :leaderID, :clubName, :bookID)"
)

func (sql *BookClubMysql) GetListClubs() (*models.ListClubs, error) {

	stmt, err := sql.db.db.Preparex(GET_CLUBS_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while preparing query for getting list of clubs: %s", err)
		return nil, err
	}

	res, err := stmt.Queryx()

	if err != nil {
		log.Printf("error while querying db for list of clubs: %s", err)
		return nil, err
	}
	defer res.Close()

	clubsList := make([]models.ClubData, 0)
	for res.Next() {

		var club models.ClubData
		err := res.StructScan(&club)

		if err != nil {
			log.Printf("error while scanning result for list of clubs: %s", err)
			return nil, err
		}

		clubsList = append(clubsList, club)

	}

	return &models.ListClubs{Clubs: clubsList}, nil

}

func (sql *BookClubMysql) GetUserClubMembers(userID string) ([]models.ClubMemberData, error) {

	stmt, err := sql.db.db.Preparex(GET_USER_CLUB_MEMBERS_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while preparing query for getting list of users clubs: %s", err)
		return nil, err
	}

	res, err := stmt.Queryx(userID)

	if err != nil {
		log.Printf("error while querying db for list of users clubs: %s", err)
		return nil, err
	}
	defer res.Close()

	clubMembersList := make([]models.ClubMemberData, 0)
	for res.Next() {

		var clubMember models.ClubMemberData
		err := res.StructScan(&clubMember)

		if err != nil {
			log.Printf("error while scanning result for list of clubs: %s", err)
			return nil, err
		}

		clubMembersList = append(clubMembersList, clubMember)

	}

	return clubMembersList, nil

}

func (sql *BookClubMysql) GetUserClubs(memberEntries []models.ClubMemberData) (*models.ListClubs, error) {

	clubsList := make([]models.ClubData, 0)

	for memberEntryID := range memberEntries {

		stmt, err := sql.db.db.Preparex(GET_CLUB_DATA_QUERY)

		if err != nil {
			log.Printf("error while querying db for club data: %s", err)
			return nil, err
		}

		row := stmt.QueryRowx(memberEntryID)
		var clubData models.ClubData
		if err = row.StructScan(&clubData); err != nil {
			log.Printf("error while scanning result for club data: %s", err)
			return nil, err
		}

		clubsList = append(clubsList, clubData)
		stmt.Close()
	}
	return &models.ListClubs{Clubs: clubsList}, nil

}

func (sql *BookClubMysql) GetClubDataForEntryID(entryID string) (*models.ClubData, error) {

	stmt, err := sql.db.db.Preparex(GET_CLUB_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while querying db for club data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(entryID)
	var clubData models.ClubData
	if err = row.StructScan(&clubData); err != nil {
		log.Printf("error while scanning result for club data: %s", err)
		return nil, err
	}

	return &clubData, nil

}

func (sql *BookClubMysql) CreateClub(createRequest *models.CreateClubRequest) error {
	stmt, err := sql.db.db.PrepareNamed(CREATE_CLUB)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while preparing club create insert: %s", err)
		return err
	}

	result, err := stmt.Exec(createRequest)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		err = errors.New(fmt.Sprintf("club already exist: ID:%s, ClubName:%s", createRequest.EntryID, createRequest.ClubName))
		log.Print(err.Error())
		return err
	}

	return nil
}
