package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const (
	authEndpoint = "/v1/auth/login"
	userEndpoint = "/v1/user"
)

//GetLoginResponse sends request to login with the arj system as a whole
func (r *Requests) GetLoginResponse(userLoginRequest *models.UserLoginRequest) (*models.ArjAPILoginResponse, error) {
	var reqBodyBytes *bytes.Buffer = new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(userLoginRequest)

	resp, err := http.Post(r.config.ArjBackendURL+authEndpoint, "application/json", reqBodyBytes)

	if err != nil {
		log.Printf("Response error #1 --> %s", err.Error())
		return nil, err
	}

	defer closeResponse(resp)
	var decodedResponse models.ArjAPILoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		log.Printf("Response error #2 --> %s", err.Error())
		return nil, err
	}

	return &decodedResponse, nil
}

//GetUserData grabs and unmarshalls user data from a provided sso token
func (r *Requests) GetUserData(ssoToken string) (*models.ArjAPIUserDataResponse, error) {

	var bearer string = "Bearer " + ssoToken

	req, err := http.NewRequest("GET", r.config.ArjBackendURL+userEndpoint, nil)

	if err != nil {
		log.Printf("Response error #1 --> %s", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", bearer)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Response error #2 --> %s", err.Error())
		return nil, err
	}
	defer closeResponse(resp)

	var decodedResponse models.ArjAPIUserDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		if err != nil {
			log.Printf("Response error #3 --> %s", err.Error())
			return nil, err
		}
	}

	return &decodedResponse, nil
}
