package models

import "github.com/google/uuid"

type JoinClubRequest struct {
	UserID  string    `json:"user_id" db:"userID"`
	ClubID  string    `json:"club_id" db:"clubID"`
	EntryID uuid.UUID `db:"entryID"`
}

type ClubDataRequest struct {
	EntryID string `json:"entry_id"`
}

type BookDataRequest struct {
	EntryID string `json:"entry_id"`
}

type CreateClubRequest struct {
	LeaderID string    `json:"user_id" db:"leaderID"` //UserID of club creator
	ClubName string    `json:"club_name" db:"clubName"`
	EntryID  uuid.UUID `db:"entryID"`
	BookID   string    `json:"book_id" db:"bookID"`
}
