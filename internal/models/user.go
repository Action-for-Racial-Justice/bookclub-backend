package models

type UserRequest struct {
	ID string `json:"id"`
}
type User struct {
	ID           string `json:"id" db:"ID"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned uint32 `json:"club_assigned" db:"clubAssigned"`
}
