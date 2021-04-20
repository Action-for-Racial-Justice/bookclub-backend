package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (sql *BookClubMysql) GetBookDataForEntryID(entryID string) (*models.BookResponse, error) {

	stmt, err := sql.db.db.Preparex(GET_BOOK_DATA_QUERY)
	defer stmt.Close()

	if err != nil {
		log.Printf("error while querying db for book data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(entryID)
	var bookData models.BookResponse
	if err = row.StructScan(&bookData); err != nil {
		log.Printf("error while scanning result for book data: %s", err)
		return nil, err
	}

	return &bookData, nil

}
