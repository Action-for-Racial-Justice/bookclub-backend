package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParameters := r.Header["id"]
	render.JSON(w, r, bh.service.GetUserData(reqParameters[0]))
}
