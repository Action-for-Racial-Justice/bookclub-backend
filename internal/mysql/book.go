package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

//GetBookDataForID returns a struct holding book fields given a supplied bookID
func (bcm *BookClubMysql) GetBookDataForID(bookID string) (*models.BookData, error) {

	stmt, err := bcm.mysql.db.Preparex(GET_BOOK_DATA_QUERY)

	defer closeStatement(stmt)
	if err != nil {
		log.Printf("error while querying db for book data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(bookID)
	var bookData models.BookData
	if err = row.StructScan(&bookData); err != nil {
		log.Printf("error while scanning result for book data: %s", err)
		return nil, err
	}

	return &bookData, nil

}
