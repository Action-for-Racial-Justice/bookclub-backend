package models

type UserDataResponse struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned string `json:"club_assigned" db:"clubAssigned"`
}

type ClubResponse struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	LeaderID string `json:"user_id" db:"leaderID"`
	ClubName string `json:"club_name" db:"clubName"`
	BookID   string `json:"book_id" db:"bookID"`
}

type ClubMemberResponse struct {
	EntryID string `json:"entry_id" db:"entryID"`
	UserID  string `json:"user_id" db:"userID"`
	ClubID  string `json:"club_id" db:"clubID"`
}

// Data for a book entry
// swagger:response BookResponse
type BookResponse struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}

type ListClubsResponse struct {
	Clubs []ClubResponse
}
