package domain

type User struct {
	Id           uint   `json:"id,omitempty"`
	Role         Role   `json:"role"`
	Username     string `json:"username,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	PasswordHash string `json:"-"`
	Salt         string `json:"-"`
}
