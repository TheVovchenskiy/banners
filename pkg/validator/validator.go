package validator

import (
	"fmt"
	"strings"

	"github.com/TheVovchenskiy/banners/pkg/utils"
)

const (
	UsernameMaxLen = 150
	PasswordMinLen = 6
)

func ValidateUsername(username string) error {
	if len(username) == 0 || len(username) > UsernameMaxLen || strings.Count(username, " ") > 0 {
		return fmt.Errorf("%w: 0 < len(username) <= %d and must not contain spaces", ErrInvalidUsername, UsernameMaxLen)
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < PasswordMinLen {
		return fmt.Errorf("%w: len(password) >= %d", ErrInvalidPassword, PasswordMinLen)
	}
	return nil
}

func ValidateRole(role string, allowedRoles []string) error {
	if !utils.In(role, allowedRoles) {
		return fmt.Errorf("%w: given role must be one of the %v", ErrInvalidRole, allowedRoles)
	}
	return nil
}

func ValidateAdminKey(role string, adminKey string, serverAdminKey string) error {
	if role == "admin" && adminKey != serverAdminKey {
		return ErrInvalidAdminKey
	}
	return nil
}
