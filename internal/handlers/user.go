package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (bh *BookClubHandler) GetUserData(w http.ResponseWriter, r *http.Request) {

	reqParameters := mux.Vars(r)

	//validate request parameters 

	userID := reqParameters["userID"]

	
}
