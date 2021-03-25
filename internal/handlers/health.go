package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// HealthCheck handle health check
func (bh *BookClubHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	render.JSON(w, r, bh.service.CheckHealth())
}
