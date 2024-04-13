package repository

import "errors"

var (
	ErrAccountAlreadyExists = errors.New("an account with given username already exists")
	ErrNoUserFound          = errors.New("no user found")
	ErrNotInserted          = errors.New("data not inserted")
)
