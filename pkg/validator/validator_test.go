package validator

import (
	"testing"

	"github.com/TheVovchenskiy/banners/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateRegisterInput(t *testing.T) {
	tests := []struct {
		name          string
		registerInput domain.RegisterInput
		expectedErr   error
	}{
		{
			"valid input",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			nil,
		},
		{
			"empty username",
			domain.RegisterInput{
				Username: "",
				Role:     "admin",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			ErrInvalidUsername,
		},
		{
			"username with spaces",
			domain.RegisterInput{
				Username: "one two",
				Role:     "admin",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			ErrInvalidUsername,
		},
		{
			"small password",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "12345",
				AdminKey: "admin_key",
			},
			ErrInvalidPassword,
		},
		{
			"small password",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "12345",
				AdminKey: "admin_key",
			},
			ErrInvalidPassword,
		},
		{
			"user role",
			domain.RegisterInput{
				Username: "username",
				Role:     "user",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			nil,
		},
		{
			"other role",
			domain.RegisterInput{
				Username: "username",
				Role:     "other_role",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			ErrInvalidRole,
		},
		{
			"invalid admin key",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "qwerty123",
				AdminKey: "invalid_admin_key",
			},
			ErrInvalidAdminKey,
		},
		{
			"invalid admin key with user role",
			domain.RegisterInput{
				Username: "username",
				Role:     "user",
				Password: "qwerty123",
				AdminKey: "invalid_admin_key",
			},
			nil,
		},
		{
			"empty admin key with user role",
			domain.RegisterInput{
				Username: "username",
				Role:     "user",
				Password: "qwerty123",
				AdminKey: "",
			},
			nil,
		},
		{
			"empty admin key with admin role",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "qwerty123",
				AdminKey: "",
			},
			ErrInvalidAdminKey,
		},
		{
			"admin key with invalid role",
			domain.RegisterInput{
				Username: "username",
				Role:     "invalid_role",
				Password: "qwerty123",
				AdminKey: "admin_key",
			},
			ErrInvalidRole,
		},
		{
			"admin without admin_key",
			domain.RegisterInput{
				Username: "username",
				Role:     "admin",
				Password: "qwerty123",
			},
			ErrInvalidAdminKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := ValidateRegisterInput(tt.registerInput, []string{"user", "admin"}, "admin_key")

			assert.ErrorIs(t, actualErr, tt.expectedErr, "errors must match")
		})
	}
}
