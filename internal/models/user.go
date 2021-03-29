package models

type UserRequest struct {
	UserID string `json:"user_id"`
}
type User struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned uint32 `json:"club_assigned" db:"clubAssigned"`
}
