package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

func (svc *BookClubService) GetUserData(userID string) (*models.UserData, error) {

	if err := svc.validator.ValidateUserID(userID); err != nil {
		return nil, err
	}

	userData, err := svc.mysql.GetUserDataForUserID(userID)
	if err != nil {
		log.Printf("Error while retrieving user data from mysql database: %s", err)
		return nil, err
	}

	return userData, nil
}

func (svc *BookClubService) GetSSOToken(userLoginRequest *models.UserLoginRequest) (string, error) {
	log.Printf("%+v", userLoginRequest)

	arjResponse, err := svc.requests.GetLoginResponse(userLoginRequest)

	log.Printf("%+v", arjResponse)
	if err != nil {
		return "", err
	}

	if !arjResponse.Success {
		return "", bcerrors.NewError("request failed", bcerrors.InternalError)
	}

	return arjResponse.Auth["token"], nil
}

func (svc *BookClubService) FetchUserDataFromToken(SSOToken string) (*models.ArjUser, error) {
	arjResponse, err := svc.requests.GetUserData(SSOToken)

	log.Printf("%+v", arjResponse)
	if err != nil {
		return nil, err
	}

	if !arjResponse.Success {
		return nil, bcerrors.NewError("request failed", bcerrors.InternalError)
	}

	return &arjResponse.User, nil
}
