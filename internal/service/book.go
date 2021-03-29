package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (svc *BookClubService) GetBookData(id string) *models.BookData {

	bookData, err := svc.mysql.GetBookDataForID(id)
	if err != nil {
		log.Printf("Error while retrieving book data from mysql database: %s", err)
		return nil
	}

	return bookData
}
