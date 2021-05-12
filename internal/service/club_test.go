package service_test

import (
	"errors"
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	clubsExpected = &models.Clubs{
		Clubs: []models.Club{
			{
				EntryID:     "1",
				LeaderID:    "2",
				ClubName:    "3",
				BookID:      "4",
				Description: "test description",
			},
			{
				EntryID:     "5",
				LeaderID:    "6",
				ClubName:    "7",
				BookID:      "8",
				Description: "test description2",
			},
		},
	}

	clubsMembersExpected = []models.ClubMember{
		{
			EntryID: "9",
			UserID:  "13",
			ClubID:  "1",
		},
		{
			EntryID: "10",
			UserID:  "13",
			ClubID:  "5",
		},
	}

	expectedClubRequest = &models.CreateClubRequest{
		LeaderID:    "1",
		ClubName:    "Keaton Club",
		BookID:      "2",
		Description: "This is the test Keaton Club",
	}
)

func TestJoinClub(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetUserDataForUserID("42").Return(nil, nil).Times(1)
	ts.mockMysql.EXPECT().CreateUserClubMember(gomock.Any()).Return(nil).Times(1)
	id, err := ts.svc.UserJoinClub(&models.JoinClubRequest{UserID: "42", ClubID: "35"})

	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestJoinClubForIDFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("GetUserDataForUserID error")

	ts.mockMysql.EXPECT().GetUserDataForUserID("").Return(nil, errorMessage).Times(1)
	id, err := ts.svc.UserJoinClub(&models.JoinClubRequest{UserID: "", ClubID: ""})

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Error(t, err)
	assert.Empty(t, id)
}

func TestJoinClubCreateUserFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("CreateUserClubMember error")

	ts.mockMysql.EXPECT().GetUserDataForUserID("42").Return(nil, nil).Times(1)
	ts.mockMysql.EXPECT().CreateUserClubMember(gomock.Any()).Return(errorMessage).Times(1)
	id, err := ts.svc.UserJoinClub(&models.JoinClubRequest{UserID: "42", ClubID: ""})

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestLeaveClubLeader(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().IsUserClubLeader(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(true, nil).Times(1)
	ts.mockMysql.EXPECT().DeleteClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(nil).Times(1)
	err := ts.svc.UserLeaveClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"})

	assert.NoError(t, err)
}

func TestLeaveClubMember(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().IsUserClubLeader(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(false, nil).Times(1)
	ts.mockMysql.EXPECT().DeleteUserClubMember(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(nil).Times(1)
	err := ts.svc.UserLeaveClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"})

	assert.NoError(t, err)
}

func TestLeaveClubIsLeaderFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("IsUserLeader error")

	ts.mockMysql.EXPECT().IsUserClubLeader(&models.LeaveClubRequest{UserID: "", ClubID: ""}).Return(false, errorMessage).Times(1)
	err := ts.svc.UserLeaveClub(&models.LeaveClubRequest{UserID: "", ClubID: ""})

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Error(t, err)
}

func TestLeaveClubDeleteClubFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("DeleteClub error")

	ts.mockMysql.EXPECT().IsUserClubLeader(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(true, nil).Times(1)
	ts.mockMysql.EXPECT().DeleteClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(errorMessage).Times(1)
	err := ts.svc.UserLeaveClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"})

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Error(t, err)
}

func TestLeaveClubDeleteMemberFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("DeleteClubClubMember error")

	ts.mockMysql.EXPECT().IsUserClubLeader(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(false, nil).Times(1)
	ts.mockMysql.EXPECT().DeleteUserClubMember(&models.LeaveClubRequest{UserID: "5", ClubID: "5"}).Return(errorMessage).Times(1)
	err := ts.svc.UserLeaveClub(&models.LeaveClubRequest{UserID: "5", ClubID: "5"})

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Error(t, err)
}

func TestGetClubData(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetClubDataForEntryID("13").Return(&clubsExpected.Clubs[0], nil).Times(1)
	club := ts.svc.GetClubData("13")

	assert.Equal(t, &clubsExpected.Clubs[0], club)
}

func TestGetClubDataForEntryIDFail(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetClubDataForEntryID("13").Return(&models.Club{
		EntryID: "not1", LeaderID: "2", ClubName: "3", BookID: "4", Description: "test description"}, nil).Times(1)
	club := ts.svc.GetClubData("13")

	assert.NotEqual(t, clubsExpected.Clubs[0], club)
}

func TestGetClubs(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetListClubs().Return(clubsExpected, nil).Times(1)
	clubs := ts.svc.GetClubs()

	assert.Equal(t, clubsExpected, clubs)
}

func TestGetUserClubs(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().GetUserClubMembers("13").Return(clubsMembersExpected, nil).Times(1)
	ts.mockMysql.EXPECT().GetUserClubs(clubsMembersExpected).Return(clubsExpected, nil).Times(1)
	clubs, err := ts.svc.GetUserClubs("13")

	assert.Equal(t, clubsExpected, clubs)
	assert.NoError(t, err)
}

func TestGetUserClubsGetMembersFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("GetUserClubMembers error")

	ts.mockMysql.EXPECT().GetUserClubMembers("12").Return(nil, errorMessage).Times(1)
	clubs, err := ts.svc.GetUserClubs("12")

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.NotEqual(t, clubsExpected, clubs)
	assert.Error(t, err)
}

func TestGetUserClubsGetUserClubsFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("GetUserClubs error")

	ts.mockMysql.EXPECT().GetUserClubMembers("").Return(clubsMembersExpected, nil).Times(1)
	ts.mockMysql.EXPECT().GetUserClubs(clubsMembersExpected).Return(nil, errorMessage).Times(1)
	clubs, err := ts.svc.GetUserClubs("")

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.NotEqual(t, clubsExpected, clubs)
	assert.Error(t, err)
}

func TestCreateClub(t *testing.T) {
	ts := createTestSuite(t)

	ts.mockMysql.EXPECT().CreateClub(gomock.Any()).Return(nil).Times(1)
	ts.mockMysql.EXPECT().UserJoinClub(gomock.Any()).Return(gomock.Any().String(), nil).Times(1)
	id, err := ts.svc.CreateClub(expectedClubRequest)

	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestCreateClubFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("CreateClubFail error")

	ts.mockMysql.EXPECT().CreateClub(gomock.Any()).Return(errorMessage).Times(1)
	id, err := ts.svc.CreateClub(expectedClubRequest)

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestCreateClubUserJoinFail(t *testing.T) {
	ts := createTestSuite(t)
	errorMessage := errors.New("CreateClubUserJoin error")

	ts.mockMysql.EXPECT().CreateClub(gomock.Any()).Return(nil).Times(1)
	ts.mockMysql.EXPECT().UserJoinClub(gomock.Any()).Return("", errorMessage).Times(1)
	id, err := ts.svc.CreateClub(expectedClubRequest)

	assert.Equal(t, errorMessage.Error(), err.Error())
	assert.Empty(t, id)
	assert.Error(t, err)
}
