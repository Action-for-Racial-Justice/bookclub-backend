package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	userResponse, err := bh.service.GetUserData(userRequest.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(err))
		return
	}

	render.JSON(w, r, userResponse)
}

func (bh *BookClubHandler) GetSSOToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signInRequest models.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&signInRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	authToken, err := bh.service.GetSSOToken(&signInRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, curateJSONError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, models.SingleSignOn{Token: authToken})
}

func (bh *BookClubHandler) GetArjBackendUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tokenizedRequest models.SingleSignOn

	if err := json.NewDecoder(r.Body).Decode(&tokenizedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	userData, err := bh.service.FetchUserDataFromToken(tokenizedRequest.Token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, curateJSONError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, userData)
}

