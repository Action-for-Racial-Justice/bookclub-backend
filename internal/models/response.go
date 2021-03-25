package models

import (
	"time"
)

type UserData struct {
	ID           int    `json:"id" db:"id"`
	FullName     string `json:"fullName" db:"fullName"`
	ClubAssigned string `json:"clubAssigned" db:"clubAssigned"`
}

//ex
//{"club": "keaton_club", "success": "true"}
type UserGetResponse struct {
	Club    *BookClub `json:"club"`
	Success bool      `json:"success"`
}

type BookClub struct {
	Name     string    `json:"name"`
	Book     string    `json:"book"`
	Deadline time.Time `json:"deadline"`
}
