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
