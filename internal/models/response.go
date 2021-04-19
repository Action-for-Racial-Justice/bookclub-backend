package models

//UserData ...
type UserData struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"fullName" db:"fullName"`
	ClubAssigned string `json:"clubAssigned" db:"clubAssigned"`
}

//UserGetResponse ...
type UserGetResponse struct {
	Club    *ClubData `json:"club"`
	Success bool      `json:"success"`
}

//ClubData ...
type ClubData struct {
	ID       string `json:"id" db:"id"`
	LeaderID string `json:"leader_id" db:"leaderId"`
	ClubName string `json:"club_name" db:"clubName"`
	BookID   string `json:"book_id" db:"bookId"`
}

//BookData ...
type BookData struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	IsActive bool   `json:"is_active" db:"isActive"`
}
