package models

import "github.com/google/uuid"

//ClubMember ...
type ClubMember struct {
	ID     string `json:"id" db:"id"`
	Uid    string `json:"uid" db:"uid"`
	ClubID uint32 `json:"club_id" db:"clubId"`
}

//JoinClubRequest ...
type JoinClubRequest struct {
	UserID string    `json:"user_id" db:"userID"`
	ClubID string    `json:"club_id" db:"clubID"`
	ID     uuid.UUID `db:"ID"`
}

//ClubRequest ...
type ClubRequest struct {
	ID string `json:"id"`
}

//BookRequest ...
type BookRequest struct {
	ID string `json:"id"`
}
