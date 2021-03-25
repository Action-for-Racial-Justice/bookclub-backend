package models

import (
	"time"
)

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
