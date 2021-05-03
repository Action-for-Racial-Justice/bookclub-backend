package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// swagger:route GET /health health getHealth
// Returns a healthy response if API dependencies are up
// responses:
//	200: HealthCheck
//	400: ErrorResponse

// HealthCheck returns a renders health check struct that snapshots dependency health
func (bh *BookClubHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	render.JSON(w, r, bh.service.CheckHealth())
}
