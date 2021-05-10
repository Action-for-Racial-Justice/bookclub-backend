package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

// swagger:route GET /book book getBook
// Returns data for a single book
// responses:
//	200: Book
//	400: ErrorResponse

// GetBookData returns the data for a single book given a BookDataRequest
func (bh *BookClubHandler) GetBookData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bookRequest models.BookDataRequest

	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}
	render.JSON(w, r, bh.service.GetBookData(bookRequest.EntryID))
}
