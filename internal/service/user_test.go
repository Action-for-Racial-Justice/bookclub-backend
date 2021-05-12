package service_test

import (
	"errors"
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	userData = &models.UserData{
		ID:       "1",
		FullName: "A Farewell to Arms",
	}

	expectedUserLoginRequest = &models.UserLoginRequest{
		Email:    "k.currie@qa.arj.today",
		Password: "Mnbvcxz1",
	}

	expectedArjAPILoginResponse = &models.ArjAPILoginResponse{
		Success: true,
	}

	expectedArjAPIUserDataResponse = &models.ArjAPIUserDataResponse{
		Success: true,
	}
)

func TestGetSSOToken(t *testing.T) {
	ts := createTestSuite(t)
	expectedArjAPILoginResponse.Auth = make(map[string]string)
	expectedArjAPILoginResponse.Auth["token"] = gomock.Any().String()

	ts.mockRequests.EXPECT().GetLoginResponse(expectedUserLoginRequest).Return(expectedArjAPILoginResponse, nil).Times(1)
	token, err := ts.svc.GetSSOToken(expectedUserLoginRequest)

	assert.NotNil(t, token)
	assert.NoError(t, err)
}

func TestGetSSOTokenFail(t *testing.T) {
	ts := createTestSuite(t)
	expectedArjAPILoginResponse.Auth = make(map[string]string)
	errorMessage := errors.New("TestGetSSOTokenFail error")

	ts.mockRequests.EXPECT().GetLoginResponse(expectedUserLoginRequest).Return(nil, errorMessage).Times(1)
	token, err := ts.svc.GetSSOToken(expectedUserLoginRequest)

	assert.Empty(t, token)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errorMessage.Error())

}

func TestFetchUserDataFromToken(t *testing.T) {
	ts := createTestSuite(t)
	expectedArjAPIUserDataResponse.User = &models.ArjUser
	expectedArjAPILoginResponse.Auth["token"] = gomock.Any().String()

	ts.mockRequests.EXPECT().GetLoginResponse(expectedUserLoginRequest).Return(expectedArjAPILoginResponse, nil).Times(1)
	token, err := ts.svc.GetSSOToken(expectedUserLoginRequest)

	assert.NotNil(t, token)
	assert.NoError(t, err)
}
