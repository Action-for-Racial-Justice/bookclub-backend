package mysql

import (
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	getAllClubsQuery        = "SELECT * FROM club;"
	getUserClubMembersQuery = "SELECT * FROM club_member where userID = ?;"
	getClubDataQuery        = "SELECT * FROM club where entryID = ?;"
	createClubQuery         = "INSERT IGNORE INTO club(entryID, leaderID, clubName, bookID, description) VALUES(:entryID, :leaderID, :clubName, :bookID, :description);"
	deleteClubMembersQuery  = "DELETE FROM club_member WHERE clubID = :clubID;"
	deleteClubQuery         = "DELETE FROM club WHERE leaderID = :userID AND entryID = :clubID;"
	addClubBookQuery        = "UPDATE club SET bookID = :bookID WHERE entryID = :clubID;"
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

//DeleteClub deletes all club member entries, and then deletes the club entry
func (bcm *BookClubMysql) DeleteClub(deleteRequest *models.LeaveClubRequest) error {

	stmt, err := bcm.mysql.db.PrepareNamed(deleteClubMembersQuery)

	if err != nil {
		log.Printf("error while preparing club members delete: %s", err)
		return err
	}

	result, err := stmt.Exec(deleteRequest)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return bcerrors.NewError(
			"club does not exist, could not delete member entries",
			bcerrors.InternalError,
		)
	}

	defer closeNamedStatement(stmt)

	stmt2, err := bcm.mysql.db.PrepareNamed(deleteClubQuery)

	if err != nil {
		log.Printf("error while preparing club delete: %s", err)
		return err
	}
	defer closeNamedStatement(stmt2)

	result2, err := stmt2.Exec(deleteRequest)
	if err != nil {
		return err
	}
	rowsAffected2, err := result2.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected2 == 0 {
		return bcerrors.NewError(
			"club does not exist, could not be deleted",
			bcerrors.InternalError,
		)
	}

	return nil
}

//AddClubBook
func (bcm *BookClubMysql) AddClubBook(addBookRequest *models.AddBookRequest) error {
	stmt, err := bcm.mysql.db.PrepareNamed(addClubBookQuery)

	if err != nil {
		log.Printf("error while preparing add book to club query: %s", err)
		return err
	}
	defer closeNamedStatement(stmt)

	result, err := stmt.Exec(addBookRequest)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return bcerrors.NewError(
			fmt.Sprintf("add club book update failed: BookID:%s, ClubID:%s",
				addBookRequest.BookID,
				addBookRequest.ClubID),
			bcerrors.InternalError,
		)
	}

	return nil
}
