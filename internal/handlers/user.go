package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/go-chi/render"
)

// swagger:route GET /user/id user getUserData
// Returns data for a user given a UserRequest
// responses:
//	200: UserData
//	400: ErrorResponse

//GetUserData retrieves user data
func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	userResponse, err := bh.service.GetUserData(userRequest.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(err))
		return
	}

	render.JSON(w, r, userResponse)
}

// swagger:route GET /user/clubs user listUserClubs
// Returns a list of clubs for a given user, given a UserClubsRequest
// responses:
//	200: Clubs
//	400: ErrorResponse

//GetUserClubs returns a list of clubs for a given user, given a UserClubsRequest
func (bh *BookClubHandler) GetUserClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userClubsRequest models.UserClubsRequest

	if err := json.NewDecoder(r.Body).Decode(&userClubsRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
	}
	clubs, err := bh.service.GetUserClubs(userClubsRequest.UserID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, clubs)
}

// swagger:route GET /user user getUserSSOToken
// Returns a sso token if exists for a email and password
// responses:
//	200:
//	400:

//GetSSOToken grabs sso token for login info
func (bh *BookClubHandler) GetSSOToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signInRequest models.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&signInRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	authToken, err := bh.service.GetSSOToken(&signInRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, curateJSONError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, models.SingleSignOn{Token: authToken})
}

// swagger:route GET /user user getUserSSOToken
// Returns a sso token if exists for a email and password
// responses:
//	200: ArjUser
//	400: ErrorResponse
//  500: ErrorResponse

//GetArjBackendUserData gets user data from ARJ monolithic api through SSO token
func (bh *BookClubHandler) GetArjBackendUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tokenizedRequest models.SingleSignOn

	if err := json.NewDecoder(r.Body).Decode(&tokenizedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	userData, err := bh.service.FetchUserDataFromToken(tokenizedRequest.Token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, curateJSONError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, userData)
}

// swagger:route DELETE /user/session user EndUserSession
// Returns Nothing
// responses:
//	200:
//	400: ErrorResponse

//EndUserSession ends user session by talking to crappy monolith API
func (bh *BookClubHandler) EndUserSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	splitToken := strings.TrimSpace(
		strings.Split(
			r.Header.Get("Authorization"), "Bearer")[1],
	)

	if err := bh.service.DeleteUserSession(splitToken); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, curateJSONError(err))
	}
	w.WriteHeader(http.StatusOK)
}
