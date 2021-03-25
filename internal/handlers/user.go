package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParameters := mux.Vars(r)
	render.JSON(w, r, bh.service.GetUserData(reqParameters["id"]))
}
