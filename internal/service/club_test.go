package service_test

import (
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

var (
	expectedJoinClub = &models.JoinClubRequest{
		UserID: "42",
		ClubID: "35",
	}
)

func TestUserJoinClub(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetUserDataForUserID("42").Return(nil, nil).Times(1)
	ts.mockMysql.EXPECT().CreateUserClubMember(expectedJoinClub).Return(nil).Times(1)

	id, err := ts.svc.UserJoinClub(&models.JoinClubRequest{UserID: "42", ClubID: "42"})

	assert.Error(t, err)
	assert.NotNil(t, id)

}
