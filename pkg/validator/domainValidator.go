package validator

import "github.com/TheVovchenskiy/banners/internal/domain"

func ValidateRegisterInput(registerInput domain.RegisterInput, allowedRoles []string, serverAdminKey string) (err error) {
	if err = ValidateUsername(registerInput.Username); err != nil {
		return
	}
	if err = ValidatePassword(registerInput.Password); err != nil {
		return
	}
	if err = ValidateRole(registerInput.Role, allowedRoles); err != nil {
		return
	}
	if err = ValidateAdminKey(registerInput.Role, registerInput.AdminKey, serverAdminKey); err != nil {
		return
	}

	return
}
