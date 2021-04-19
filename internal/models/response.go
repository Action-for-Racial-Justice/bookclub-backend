package models

type UserData struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned string `json:"club_assigned" db:"clubAssigned"`
}

//ex
//{"club": "keaton_club", "success": "true"}
type UserGetResponse struct {
	Club    *ClubData `json:"club"`
	Success bool      `json:"success"`
}

type ClubData struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	LeaderID string `json:"user_id" db:"leaderID"`
	ClubName string `json:"club_name" db:"clubName"`
	BookID   string `json:"book_id" db:"bookID"`
}

type ClubMemberData struct {
	EntryID string `json:"entry_id" db:"entryID"`
	UserID  string `json:"user_id" db:"userID"`
	ClubID  string `json:"club_id" db:"clubID"`
}

type BookData struct {
	EntryID  string `json:"entry_id" db:"entryID"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}

type ListClubs struct {
	Clubs []ClubData
}
