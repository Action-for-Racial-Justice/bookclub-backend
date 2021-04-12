package requests

import (
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/google/wire"
)

var Module = wire.NewSet(
	New,
)

type Config struct {
	ArjBackendURL string
}

type IRequests interface {
	GetLoginResponse(userLoginRequest *models.UserLoginRequest) (*models.ArjAPILoginResponse, error)
	GetUserData(SSOToken string) (*models.ArjAPIUserDataResponse, error)
}

type Requests struct {
	config *Config
}

func New(cfg *Config) *Requests {

	return &Requests{config: cfg}
}
