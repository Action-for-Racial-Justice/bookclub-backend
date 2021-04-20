package models

// UserData is a user for the bookclub application
// swagger:model UserData

//UserData ...
type UserData struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned string `json:"club_assigned" db:"clubAssigned"`
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

// Book is a book data entry
// swagger:model Book

//Book ...
type Book struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}

// Clubs is a list of clubs
// swagger:model Clubs

//Clubs ...
type Clubs struct {
	Clubs []Club
}
