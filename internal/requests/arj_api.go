package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	authEndpoint = "/v1/auth/login"
	userEndpoint = "/v1/user"
)

//GetLoginResponse sends request to login with the arj system as a whole
func (r *Requests) GetLoginResponse(userLoginRequest *models.UserLoginRequest) (*models.ArjAPILoginResponse, error) {
	var reqBodyBytes *bytes.Buffer = new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(userLoginRequest); err != nil {
		return nil, bcerrors.NewError(
			"Could not decode user login request",
			bcerrors.DecodeError,
		).WithRootCause(err)

	}

	resp, err := http.Post(r.config.ArjBackendURL+authEndpoint, "application/json", reqBodyBytes)

	if err != nil {
		log.Printf("[GetLoginResponse] Response error #1 --> %s", err.Error())
		return nil, err
	}

	defer closeResponse(resp)

	if resp.StatusCode != 200 {
		return nil, bcerrors.NewError(fmt.Sprintf("request failed, expected status 200, got %d", resp.StatusCode), bcerrors.InternalError)
	}
	var decodedResponse models.ArjAPILoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		log.Printf("[GetLoginResponse] Response error #2 --> %s", err.Error())
		return nil, err
	}

	return &decodedResponse, nil
}

//GetUserData grabs and unmarshalls user data from a provided sso token
func (r *Requests) GetUserData(ssoToken string) (*models.ArjAPIUserDataResponse, error) {

	var bearer string = "Bearer " + ssoToken

	req, err := http.NewRequest("GET", r.config.ArjBackendURL+userEndpoint, nil)

	if err != nil {
		log.Printf("[GetUserData] Response error #1 --> %s", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", bearer)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[GetUserData] Response error --> %s", err.Error())
		return nil, err
	}
	defer closeResponse(resp)

	var decodedResponse models.ArjAPIUserDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		if err != nil {
			log.Printf("[GetUserData] Response error #3 --> %s", err.Error())
			return nil, err
		}
	}

	return &decodedResponse, nil
}

//EndUserSession ends user session by sending request to crummy backend api
func (r *Requests) EndUserSession(ssoToken string) error {
	var bearer string = "Bearer " + ssoToken

	req, err := http.NewRequest("DELETE", r.config.ArjBackendURL+authEndpoint, nil)

	if err != nil {
		log.Printf("[EndUserSession] Response error #1 --> %s", err.Error())
		return err
	}
	req.Header.Add("Authorization", bearer)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[EndUserSession] Response error #2 --> %s", err.Error())
		return err
	}
	defer closeResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return bcerrors.NewError("Delete request failed", resp.StatusCode)
	}
	return nil
}
