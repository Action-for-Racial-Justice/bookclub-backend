package service

import (
	"time"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/google/wire"
)

type Service interface {
	CheckHealth() *models.HealthCheck
}

var Module = wire.NewSet(
	New,
)

type BookClubService struct {
}

func New() *BookClubService {
	return &BookClubService{}
}

func (svc *BookClubService) CheckHealth() *models.HealthCheck {
	return &models.HealthCheck{Timestamp: time.Now(), Healthy: true}
}
