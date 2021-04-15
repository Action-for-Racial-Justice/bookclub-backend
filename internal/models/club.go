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

type ClubDataRequest struct {
	ID string `json:"id"`
}

type BookDataRequest struct {
	ID string `json:"id"`
}

type CreateClubRequest struct {
	LeaderID string    `json:"leader_id" db:"leaderID"` //UserID of club creator
	ClubName string    `json:"club_name" db:"clubName"`
	ID       uuid.UUID `db:"ID"`
	BookID   string    `json:"book_id" db:"bookID"`
}
