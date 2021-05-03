package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
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
		return
	}
	render.JSON(w, r, bh.service.GetBookData(bookRequest.EntryID))
}

func (bh *BookClubHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query, found := r.URL.Query()["query"]

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(
			bcerrors.NewError("Query parameter not supplied", bcerrors.ValidationError),
		))
		return
	}

	bookItems, err := bh.service.SearchBooks(query[0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, bookItems)
}

// func (bh *BookClubHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// }
