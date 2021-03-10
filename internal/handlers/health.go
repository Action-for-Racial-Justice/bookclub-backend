package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// HealthCheck handle health check
func (bh *BookClubHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	render.JSON(w, r, bh.service.CheckHealth())
}
