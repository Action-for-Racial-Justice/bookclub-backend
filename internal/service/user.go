package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

//GetUserData returns user data for a user id
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

//GetSSOToken returns single sign on token for a user login request
func (svc *BookClubService) GetSSOToken(userLoginRequest *models.UserLoginRequest) (string, error) {
	arjResponse, err := svc.requests.GetLoginResponse(userLoginRequest)

	if err != nil {
		return "", err
	}

	if !arjResponse.Success {
		return "", bcerrors.NewError("request failed", bcerrors.InternalError)
	}

	return arjResponse.Auth["token"], nil
}

//FetchUserDataFromToken gets user data from monolith API for a provided sso session token
func (svc *BookClubService) FetchUserDataFromToken(SSOToken string) (*models.ArjUser, error) {
	arjResponse, err := svc.requests.GetUserData(SSOToken)

	if err != nil {
		return nil, err
	}

	if !arjResponse.Success {
		return nil, bcerrors.NewError("request failed", bcerrors.InternalError)
	}

	return &arjResponse.User, nil
}
