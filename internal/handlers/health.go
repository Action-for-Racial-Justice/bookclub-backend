package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// HealthCheck handle health check
func (bh *BookClubHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, bh.service.CheckHealth())
}
