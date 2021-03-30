package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_CLUB_DATA_QUERY = "SELECT * FROM club where id = ?"

func (sql *BookClubMysql) GetClubDataForID(id string) (*models.ClubData, error) {

	stmt, err := sql.db.db.Preparex(GET_CLUB_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while querying db for club data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(id)
	var clubData models.ClubData
	if err = row.StructScan(&clubData); err != nil {
		log.Printf("error while scanning result for club data: %s", err)
		return nil, err
	}

	return &clubData, nil

}
