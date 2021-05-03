package models

import "github.com/google/uuid"

//UserClubsRequest ...
type UserClubsRequest struct {
	UserID string `json:"user_id" db:"userID"`
}

// Clubs is a list of clubs
// swagger:model Clubs
//Clubs ...
type Clubs struct {
	Clubs []Club
}

//ClubDataRequest ...
type ClubDataRequest struct {
	EntryID string `json:"entry_id"`
}

// CreateClubRequest contains a timestamp and a boolean indicator
// swagger:model CreateClubRequest
//CreateClubRequest struct to describe current system health
//CreateClubRequest ...
type CreateClubRequest struct {
	LeaderID string    `json:"user_id" db:"leaderID"` //UserID of club creator
	ClubName string    `json:"club_name" db:"clubName"`
	EntryID  uuid.UUID `db:"entryID"`
	BookID   string    `json:"book_id" db:"bookID"`
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

// Club is a book club
// swagger:model Club
//Club ...
type Club struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	LeaderID string `json:"user_id" db:"leaderID"`
	ClubName string `json:"club_name" db:"clubName"`
	BookID   string `json:"book_id" db:"bookID"`
}
