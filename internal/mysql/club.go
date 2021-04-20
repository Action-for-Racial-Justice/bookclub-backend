package mysql

import (
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	getAllClubsQuery        = "SELECT * FROM club"
	getUserClubMembersQuery = "SELECT * FROM club_member where userID = ?"
	getClubDataQuery        = "SELECT * FROM club where entryID = ?"
	createClubQuery         = "INSERT INTO club(entryID, leaderID, clubName, bookID) VALUES(:entryID, :leaderID, :clubName, :bookID)"
)

//GetListClubs gets slice of all the clubs
func (bcm *BookClubMysql) GetListClubs() (*models.Clubs, error) {

	stmt, err := bcm.mysql.db.Preparex(getAllClubsQuery)
	defer closeStatement(stmt)

	if err != nil {
		log.Printf("error while preparing query for getting list of clubs: %s", err)
		return nil, err
	}

	rows, err := stmt.Queryx()

	if err != nil {
		log.Printf("error while querying db for list of clubs: %s", err)
		return nil, err
	}
	defer closeRows(rows)

	clubsList := make([]models.Club, 0)
	for rows.Next() {

		var club models.Club
		err := rows.StructScan(&club)

		if err != nil {
			log.Printf("error while scanning result for list of clubs: %s", err)
			return nil, err
		}

		clubsList = append(clubsList, club)

	}

	return &models.Clubs{Clubs: clubsList}, nil

}

//GetUserClubMembers gets all club member entries per user
func (bcm *BookClubMysql) GetUserClubMembers(userID string) ([]models.ClubMember, error) {

	stmt, err := bcm.mysql.db.Preparex(getUserClubMembersQuery)

	if err != nil {
		log.Printf("error while preparing query for getting list of users clubs: %s", err)
		return nil, err
	}
	defer closeStatement(stmt)

	rows, err := stmt.Queryx(userID)

	if err != nil {
		log.Printf("error while querying db for list of users clubs: %s", err)
		return nil, err
	}
	defer closeRows(rows)

	clubMembersList := make([]models.ClubMember, 0)
	var clubMember models.ClubMember

	for rows.Next() {

		if err := rows.StructScan(&clubMember); err != nil {
			log.Printf("error while scanning result for list of clubs: %s", err)
			return nil, err
		}
		clubMembersList = append(clubMembersList, clubMember)

	}

	return clubMembersList, nil

}

//GetUserClubs returns a slice of all the clubs for a club member entry slice
func (bcm *BookClubMysql) GetUserClubs(memberEntries []models.ClubMember) (*models.Clubs, error) {

	clubsList := make([]models.Club, len(memberEntries))
	var clubData models.Club

	for index, memberEntry := range memberEntries {

		stmt, err := bcm.mysql.db.Preparex(getClubDataQuery)

		if err != nil {
			log.Printf("error while querying db for club data: %s", err)
			return nil, err
		}
		defer closeStatement(stmt)
		if err = stmt.QueryRowx(memberEntry.ClubID).StructScan(&clubData); err != nil {
			log.Printf("error while scanning result for club data: %s", err)
			return nil, err
		}

		log.Printf("club Data: %+v", clubData)
		clubsList[index] = clubData
	}
	return &models.Clubs{Clubs: clubsList}, nil

}

//GetClubDataForEntryID gets club data for a club entry ID
func (bcm *BookClubMysql) GetClubDataForEntryID(entryID string) (*models.Club, error) {

	stmt, err := bcm.mysql.db.Preparex(getClubDataQuery)

	if err != nil {
		log.Printf("error while querying db for club data: %s", err)
		return nil, err
	}
	defer closeStatement(stmt)

	var clubData models.Club
	if err = stmt.QueryRowx(entryID).StructScan(&clubData); err != nil {
		log.Printf("error while scanning result for club data: %s", err)
		return nil, err
	}

	return &clubData, nil

}

//CreateClub inserts row into club table
func (bcm *BookClubMysql) CreateClub(createRequest *models.CreateClubRequest) error {
	stmt, err := bcm.mysql.db.PrepareNamed(createClubQuery)

	if err != nil {
		log.Printf("error while preparing club create insert: %s", err)
		return err
	}
	defer closeNamedStatement(stmt)

	result, err := stmt.Exec(createRequest)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return bcerrors.NewError(
			fmt.Sprintf("club already exist: ID:%s, ClubName:%s",
				createRequest.EntryID,
				createRequest.ClubName),
			bcerrors.InternalError,
		)
	}

	return nil
}
