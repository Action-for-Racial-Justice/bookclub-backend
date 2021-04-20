package models

//UserRequest ...
type UserRequest struct {
	ID string `json:"id"`
}

//User ...
type User struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"fullName"`
	ClubAssigned uint32 `json:"club_assigned" db:"clubAssigned"`
}

//UserLoginRequest ...
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//ArjAPILoginResponse ...
type ArjAPILoginResponse struct {
	Auth    map[string]string `json:"auth"`
	Success bool              `json:"success"`
}

//ArjUser ...
type ArjUser struct {
	GUID             string `json:"guid"`
	FullName         string `json:"fullName"`
	Email            string `json:"email"`
	Locale           string `json:"locale"`
	Timezone         string `json:"timezone"`
	ZipCode          string `json:"zip"`
	NickName         string `json:"nickname"`
	Race             string `json:"race"`
	Address1         string `json:"address1"`
	Address2         string `json:"address2"`
	City             string `json:"city"`
	State            string `json:"state"`
	Phone            string `json:"phone"`
	EmailVerified    bool   `json:"emailVerified"`
	IsAdmin          bool   `json:"isAdmin"`
	IsFacilitator    bool   `json:"isFacilitator"`
	IsObserver       bool   `json:"isObserver"`
	IsEnvoy          bool   `json:"isEnvoy"`
	IsFlagged        bool   `json:"isFlagged"`
	PodRestrictLocal bool   `json:"podRestrictLocal"`
	CreatedAt        string `json:"createdAt"`
	LastModified     string `json:"lastModified"`
}

// ArjAPIUserDataResponse ...
type ArjAPIUserDataResponse struct {
	Success bool    `json:"success"`
	User    ArjUser `json:"user"`
}

//SingleSignOn ...
type SingleSignOn struct {
	Token string `json:"sso_token"`
}
