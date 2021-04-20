package mysql

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	getBookDataQuery = "SELECT * FROM book where entryID = ?"
)

//GetBookDataForEntryID returns book data struct for a provided book entry ID
func (bcm *BookClubMysql) GetBookDataForEntryID(entryID string) (*models.Book, error) {

	stmt, err := bcm.mysql.db.Preparex(getBookDataQuery)

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
