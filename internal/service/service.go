//go:generate mockgen -package=mocks -destination=../mocks/service.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/service Service

package service

import (
	"time"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/requests"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/validator"

	"github.com/google/wire"
)

//Service interface to describe BookClubService struct receiver functions
type Service interface {
	AddClubBook(addBookRequest *models.AddBookRequest) error
	CheckHealth() *models.HealthCheck
	CreateClub(joinRequest *models.CreateClubRequest) (string, error)
	FetchUserDataFromToken(string) (*models.ArjUser, error)
	GetBookData(string) *models.Book
	GetClubData(string) *models.Club
	GetClubs() *models.Clubs
	GetSSOToken(userLoginRequest *models.UserLoginRequest) (string, error)
	GetUserClubs(string) (*models.Clubs, error)
	GetUserData(string) (*models.UserData, error)
	UserJoinClub(joinRequest *models.JoinClubRequest) (string, error)
	UserLeaveClub(leaveRequest *models.LeaveClubRequest) error
}

//Module to denote wire binding function
var Module = wire.NewSet(
	New,
)

//BookClubService struct to hold relevant inner data members and hold functions for business logic
type BookClubService struct {
	requests  requests.IRequests
	validator validator.Validator
	mysql     mysql.Mysql
}

//New ... constructor
func New(db mysql.Mysql, requests requests.IRequests, validator validator.Validator) *BookClubService {
	return &BookClubService{
		mysql:     db,
		requests:  requests,
		validator: validator,
	}
}

//CheckHealth checks API dependencies and returns health check struct accordingly
func (svc *BookClubService) CheckHealth() *models.HealthCheck {
	return &models.HealthCheck{Timestamp: time.Now(), Healthy: true}
}
