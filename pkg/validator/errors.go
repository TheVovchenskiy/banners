package validator

import "errors"

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidAdminKey = errors.New("invalid admin key")
)
