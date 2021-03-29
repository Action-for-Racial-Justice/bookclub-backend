//go:generate mockgen -package=mocks -destination=../mocks/service.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/service Service

package service

import (
	"time"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/google/wire"
)

//Service interface to describe BookClubService struct receiver functions
type Service interface {
	CheckHealth() *models.HealthCheck
	GetUserData(string) *models.UserData
	GetClubData(string) *models.ClubData
	GetBookData(string) *models.BookData
	UserJoinClub(joinRequest *models.JoinClubRequest) (string, error)
}

//Module to denote wire binding function
var Module = wire.NewSet(
	New,
)

//BookClubService struct to hold relevant inner data members and hold functions for business logic
type BookClubService struct {
	mysql mysql.Mysql
}

//New ... constructor
func New(db mysql.Mysql) *BookClubService {
	return &BookClubService{
		mysql: db,
	}
}

//CheckHealth checks API dependencies and returns health check struct accordingly
func (svc *BookClubService) CheckHealth() *models.HealthCheck {
	return &models.HealthCheck{Timestamp: time.Now(), Healthy: true}
}
