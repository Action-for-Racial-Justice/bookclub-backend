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
	expectedArjAPIUserDataResponse.User = models.ArjUser{GUID: "2", FullName: "Keaton"}

	ts.mockRequests.EXPECT().GetUserData("2").Return(expectedArjAPIUserDataResponse, nil).Times(1)
	arjUser, err := ts.svc.FetchUserDataFromToken("2")

	assert.NotNil(t, arjUser)
	assert.NoError(t, err)
}

func TestFetchUserDataFromTokenFail(t *testing.T) {
	ts := createTestSuite(t)
	expectedArjAPIUserDataResponse.User = models.ArjUser{GUID: "2", FullName: "Keaton"}
	errorMessage := errors.New("TestFetchUserDataFromTokenFail error")

	ts.mockRequests.EXPECT().GetUserData("2").Return(nil, errorMessage).Times(1)
	arjUser, err := ts.svc.FetchUserDataFromToken("2")

	assert.Empty(t, arjUser)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errorMessage.Error())
}

func TestDeleteUserSession(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockRequests.EXPECT().EndUserSession("2").Return(nil).Times(1)
	err := ts.svc.DeleteUserSession("2")

	assert.NoError(t, err)
}

func TestDeleteUserSessionFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("TestDeleteUserSessionFail error")

	ts.mockRequests.EXPECT().EndUserSession("2").Return(errorMessage).Times(1)
	err := ts.svc.DeleteUserSession("2")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), errorMessage.Error())
}
