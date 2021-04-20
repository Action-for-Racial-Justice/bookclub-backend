//go:generate mockgen -package=mocks -destination=../mocks/requests.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/requests IRequests

package requests

import (
	"log"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/google/wire"
)

//Module for wire binding
var Module = wire.NewSet(
	New,
)

//Config ...
type Config struct {
	ArjBackendURL string
}

//IRequests to hold request functions
type IRequests interface {
	GetLoginResponse(*models.UserLoginRequest) (*models.ArjAPILoginResponse, error)
	GetUserData(string) (*models.ArjAPIUserDataResponse, error)
}

//Requests ...
type Requests struct {
	config *Config
}

//New .. constructor
func New(cfg *Config) *Requests {

	return &Requests{config: cfg}
}

func closeResponse(response *http.Response) {
	if err := response.Body.Close(); err != nil {
		log.Printf("%s", err.Error())
	}
}
