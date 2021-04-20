package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (bcm *BookClubMysql) GetBookDataForEntryID(entryID string) (*models.Book, error) {

	stmt, err := bcm.mysql.db.Preparex(GET_BOOK_DATA_QUERY)

	defer closeStatement(stmt)
	if err != nil {
		log.Printf("error while querying db for book data: %s", err)
		return nil, err
	}

	var bookData models.Book
	if err = stmt.QueryRowx(entryID).StructScan(&bookData); err != nil {
		log.Printf("error while scanning result for book data: %s", err)
		return nil, err
	}

	return &bookData, nil

}
