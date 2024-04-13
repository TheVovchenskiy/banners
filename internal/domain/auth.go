package domain

import (
	"strings"

	"github.com/TheVovchenskiy/banners/pkg/hash"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	AdminKey string `json:"adminKey,omitempty"`
}

func (i *RegisterInput) Trim() {
	i.Username = strings.TrimSpace(i.Username)
	i.Role = strings.TrimSpace(i.Role)
}

func (i *RegisterInput) ToUser(salt string) *User {
	user := User{
		Role:         Role{Name: i.Role},
		Username:     i.Username,
		PasswordHash: hash.HashPassword(i.Password, salt),
		Salt:         salt,
	}

	return &user
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *LoginInput) Trim() {
	i.Username = strings.TrimSpace(i.Username)
}
