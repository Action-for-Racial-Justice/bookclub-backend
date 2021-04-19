package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
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

	listClubs := make([]models.ClubData, 0)
	for res.Next() {

		var club models.ClubData
		err := res.StructScan(&club)

		if err != nil {
			log.Printf("error while scanning result for list of clubs: %s", err)
		}

		listClubs = append(listClubs, club)

	}

	return &models.ListClubs{Clubs: listClubs}, nil

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
