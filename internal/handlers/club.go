package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

func (bh *BookClubHandler) CreateUserClubMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userJoinRequest models.JoinClubRequest

	if err := json.NewDecoder(r.Body).Decode(&userJoinRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	confirmationID, err := bh.service.UserJoinClub(&userJoinRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, confirmationID)
}

func (bh *BookClubHandler) GetClubData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var clubRequest models.ClubRequest

	if err := json.NewDecoder(r.Body).Decode(&clubRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}
	render.JSON(w, r, bh.service.GetClubData(clubRequest.ID))
}
