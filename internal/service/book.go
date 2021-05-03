package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

//GetBookData returns book struct for a provided book ID
func (svc *BookClubService) GetBookData(id string) *models.Book {

	bookData, err := svc.mysql.GetBookDataForEntryID(id)
	if err != nil {
		log.Printf("Error while retrieving book data from mysql database: %s", err)
		return nil
	}

	return bookData
}

func (svc *BookClubService) SearchBooks(query string) ([]*models.BookResult, error) {

	resp, err := svc.requests.QueryBooksByName(query)
	if err != nil {
		return nil, bcerrors.NewError("Failed to get response from google books API", bcerrors.InternalError)
	}
	return resp.Items, nil
}
