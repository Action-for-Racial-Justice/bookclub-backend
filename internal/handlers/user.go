package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

// swagger:route GET /user/id user getUserData
// Returns data for a user given a UserRequest
// responses:
//	200: UserData
//	400: Error

//Returns data for a user given a UserRequest
func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	userResponse, err := bh.service.GetUserData(userRequest.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(err))
		return
	}

	render.JSON(w, r, userResponse)
}

// swagger:route GET /user/clubs user listUserClubs
// Returns a list of clubs for a given user, given a UserClubsRequest
// responses:
//	200: Clubs
//	400:

//Returns a list of clubs for a given user, given a UserClubsRequest
func (bh *BookClubHandler) GetUserClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userClubsRequest models.UserClubsRequest

	if err := json.NewDecoder(r.Body).Decode(&userClubsRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}
	clubs, err := bh.service.GetUserClubs(userClubsRequest.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, clubs)
}
