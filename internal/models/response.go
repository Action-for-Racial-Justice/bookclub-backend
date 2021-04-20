package models

// UserData is a user for the bookclub application
// swagger:model UserData
type UserData struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned string `json:"club_assigned" db:"clubAssigned"`
}

// Club is a book club
// swagger:model Club
type Club struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	LeaderID string `json:"user_id" db:"leaderID"`
	ClubName string `json:"club_name" db:"clubName"`
	BookID   string `json:"book_id" db:"bookID"`
}

// ClubMember is a record of a users membership in a given book club
// swagger:model ClubMember
type ClubMember struct {
	EntryID string `json:"entry_id" db:"entryID"`
	UserID  string `json:"user_id" db:"userID"`
	ClubID  string `json:"club_id" db:"clubID"`
}

// Book is a book data entry
// swagger:model Book
type Book struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}

// Clubs is a list of clubs
// swagger:model Clubs
type Clubs struct {
	Clubs []Club
}
