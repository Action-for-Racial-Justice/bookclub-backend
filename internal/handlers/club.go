package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

// swagger:route POST /club/join club newClubMember
// Creates a new club member entry for the user given a JoinClubRequest and response returns the ClubMember EntryID
// responses:
//	200: ClubMember EntryID
//	400: ErrorResponse

//JoinClub creates a new club member entry for the user given a JoinClubRequest
func (bh *BookClubHandler) JoinClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userJoinRequest models.JoinClubRequest

	if err := json.NewDecoder(r.Body).Decode(&userJoinRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	confirmationID, err := bh.service.UserJoinClub(&userJoinRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, confirmationID)
}

// swagger:route GET /club/id club getClubData
// Returns data for a club entry given a ClubDataRequest
// responses:
//	200: Club
//	400: ErrorResponse

//GetClubData gets data for a club entry given a ClubDataRequest
func (bh *BookClubHandler) GetClubData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var clubRequest models.ClubDataRequest

	if err := json.NewDecoder(r.Body).Decode(&clubRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}
	render.JSON(w, r, bh.service.GetClubData(clubRequest.EntryID))
}

// swagger:route GET /club club listClubs
// Returns a list of clubs
// responses:
//	200: Clubs
//	400: ErrorResponse

//GetClubs renders a list of clubs
func (bh *BookClubHandler) GetClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, bh.service.GetClubs())
}

// swagger:route POST /club/create club createClub
// Creates a new club entry for the user given a CreateClubRequest and response returns the Club EntryID
// responses:
//	200: Club EntryID
//	400: ErrorResponse

//CreateClub creates a new club entry for the user given a CreateClubRequest
func (bh *BookClubHandler) CreateClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createClubRequest models.CreateClubRequest

	if err := json.NewDecoder(r.Body).Decode(&createClubRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	entryID, err := bh.service.CreateClub(&createClubRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, entryID)
}

// swagger:route POST /club/leave club leaveClub
// Leaves a club by deleting the club member entry/entries. If the user is the leader, the club is deleted.
// responses:
//	200:
//	400: Error

//LeaveClub leaves a club, or deletes the club if the user is the leader
func (bh *BookClubHandler) LeaveClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var leaveClubRequest models.LeaveClubRequest

	if err := json.NewDecoder(r.Body).Decode(&leaveClubRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	err := bh.service.UserLeaveClub(&leaveClubRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, http.StatusOK)
}

// swagger:route POST /club/book club addClubBook
// Adds a book ID to a given club
// responses:
//	200: Club EntryID
//	400: ErrorResponse

//CreateClub creates a new club entry for the user given a CreateClubRequest
func (bh *BookClubHandler) AddClubBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var addBookRequest models.AddBookRequest

	if err := json.NewDecoder(r.Body).Decode(&addBookRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	err := bh.service.AddClubBook(&addBookRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, http.StatusOK)
}