package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_BOOK_DATA_QUERY = "SELECT * FROM book where id = ?"

func (sql *BookClubMysql) GetBookDataForID(id string) (*models.BookData, error) {

	stmt, err := sql.db.db.Preparex(GET_BOOK_DATA_QUERY)
	if err != nil {
		log.Printf("error while querying db for book data: %s", err)
		return nil, err
	}

	row := stmt.QueryRowx(id)
	var bookData models.BookData
	if err = row.StructScan(&bookData); err != nil {
		log.Printf("error while scanning result for book data: %s", err)
		return nil, err
	}

	return &bookData, nil

}
