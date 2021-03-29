package models

import "github.com/google/uuid"

type ClubMember struct {
	ID     string `json:"id" db:"id"`
	Uid    string `json:"uid" db:"uid"`
	ClubID uint32 `json:"club_id" db:"clubId"`
}

type JoinClubRequest struct {
	UserID string    `json:"user_id" db:"userID"`
	ClubID string    `json:"club_id" db:"clubID"`
	ID     uuid.UUID `db:"ID"`
}

type ClubRequest struct {
	ID string `json:"id"`
}

type BookRequest struct {
	ID string `json:"id"`
}
