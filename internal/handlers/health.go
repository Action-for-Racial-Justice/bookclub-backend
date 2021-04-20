package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// swagger:route GET /health health getHealth
// Returns a healthy response if API dependencies are up
// responses:
//	200: HealthCheck
//	400: Error

//Returns a list of clubs for a given user, given a UserClubsRequest
func (bh *BookClubHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	render.JSON(w, r, bh.service.CheckHealth())
}
