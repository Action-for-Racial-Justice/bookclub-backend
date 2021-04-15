package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/google/uuid"
)

func (svc *BookClubService) UserJoinClub(joinRequest *models.JoinClubRequest) (string, error) {

	if _, err := svc.mysql.GetUserDataForUserID(joinRequest.UserID); err != nil {
		return "", err
	}

	id := uuid.New()
	joinRequest.ID = id
	//TODO validate user struct values exist

	if err := svc.mysql.CreateUserClubMember(joinRequest); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (svc *BookClubService) GetClub(userID string) *models.ClubData {

	clubData, err := svc.mysql.GetClubForID(userID)
	if err != nil {
		log.Printf("Error while retrieving club data from mysql database: %s", err)
		return nil
	}

	return clubData
}

func (svc *BookClubService) GetClubs() *models.ListClubs {

	clubs, err := svc.mysql.GetListClubs()
	if err != nil {
		log.Printf("Error while retrieving a list of clubs from mysql database: %s", err)
		return nil
	}

	return clubs
}

// func (svc *BookClubService) UserCreateClub(userID string) *models.ClubData {

// 	userData, err := svc.mysql.GetClubDataForID(userID)
// 	if err != nil {
// 		log.Printf("Error while retrieving club data from mysql database: %s", err)
// 		return nil
// 	}

// 	return userData
// }
