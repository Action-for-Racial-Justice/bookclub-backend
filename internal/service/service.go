package service

import (
	"time"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mysql"
	"github.com/google/wire"
)

type Service interface {
	CheckHealth() *models.HealthCheck
}

var Module = wire.NewSet(
	New,
)

type BookClubService struct {
	DB *mysql.Database
}

func New(DB *mysql.Database) *BookClubService {
	return &BookClubService{
		DB: DB,
	}
}

func (svc *BookClubService) CheckHealth() *models.HealthCheck {
	return &models.HealthCheck{Timestamp: time.Now(), Healthy: true}
}
