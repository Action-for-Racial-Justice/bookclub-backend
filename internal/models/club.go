package models

import "github.com/google/uuid"

//UserClubsRequest ...
type UserClubsRequest struct {
	UserID string `json:"user_id" db:"userID"`
}

//ClubDataRequest ...
type ClubDataRequest struct {
	EntryID string `json:"entry_id"`
}

//BookDataRequest ...
type BookDataRequest struct {
	EntryID string `json:"entry_id"`
}

// CreateClubRequest contains a timestamp and a boolean indicator
// swagger:model CreateClubRequest
//CreateClubRequest struct to describe current system health
//CreateClubRequest ...
type CreateClubRequest struct {
	LeaderID    string    `json:"user_id" db:"leaderID"` //UserID of club creator
	ClubName    string    `json:"club_name" db:"clubName"`
	EntryID     uuid.UUID `db:"entryID"`
	BookID      string    `json:"book_id" db:"bookID"`
	Description string    `json:"description" db:"bookDescription"`
}

//ClubMember ...
type ClubMember struct {
	EntryID string `json:"entry_id" db:"entryID"`
	UserID  string `json:"user_id" db:"userID"`
	ClubID  string `json:"club_id" db:"clubID"`
}

//JoinClubRequest ...
type JoinClubRequest struct {
	UserID  string    `json:"user_id" db:"userID"`
	ClubID  string    `json:"club_id" db:"clubID"`
	EntryID uuid.UUID `db:"entryID"`
}

//LeaveClubRequest ...
type LeaveClubRequest struct {
	UserID string `json:"user_id" db:"userID"`
	ClubID string `json:"club_id" db:"clubID"`
}
