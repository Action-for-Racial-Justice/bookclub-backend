package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (svc *BookClubService) GetUserData(userID string) (*models.UserData, error) {

	// if err := svc.validator.ValidateUserID(userID); err != nil {
	// 	return nil, err
	// }

	userData, err := svc.mysql.GetUserDataForUserID(userID)
	if err != nil {
		log.Printf("Error while retrieving user data from mysql database: %s", err)
		return nil, err
	}

	return userData, nil
}
