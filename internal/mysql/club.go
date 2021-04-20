package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_CLUB_DATA_QUERY = "SELECT * FROM club where id = ?"

//GetClubDataForID returns struct holding club data for an associated clubID from the clubs table
func (bcm *BookClubMysql) GetClubDataForID(clubID string) (*models.ClubData, error) {

	stmt, err := bcm.mysql.db.Preparex(GET_CLUB_DATA_QUERY)
	defer closeStatement(stmt)

	if err != nil {
		log.Printf("error while querying db for club data: %s", err)
		return nil, err
	}

	var clubData models.ClubData
	if err = stmt.QueryRowx(clubID).StructScan(&clubData); err != nil {
		log.Printf("error while scanning result for club data: %s", err)
		return nil, err
	}

	return &clubData, nil

}
