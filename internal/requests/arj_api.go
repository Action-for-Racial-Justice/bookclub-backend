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
	userEndpoint = "/user"
)

//GetLoginResponse
func (r *Requests) GetLoginResponse(userLoginRequest *models.UserLoginRequest) (*models.ArjAPILoginResponse, error) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(userLoginRequest)

	resp, err := http.Post(r.config.ArjBackendURL+authEndpoint, "application/json", reqBodyBytes)

	if err != nil {
		log.Printf("Response error #1 --> %s", err.Error())
		return nil, err
	}

	println("status code", resp.StatusCode)
	defer resp.Body.Close()
	var decodedResponse models.ArjAPILoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		if err != nil {
			log.Printf("Response error #2 --> %s", err.Error())
			return nil, err
		}
	}

	return &decodedResponse, nil
}

func (r *Requests) GetUserData(SSOToken string) (*models.ArjAPIUserDataResponse, error) {

	var bearer string = "Bearer " + SSOToken

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

	var decodedResponse models.ArjAPIUserDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		if err != nil {
			log.Printf("Response error #3 --> %s", err.Error())
			return nil, err
		}
	}

	return &decodedResponse, nil
}
