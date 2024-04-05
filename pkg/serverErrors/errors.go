package serverErrors

import (
	"errors"
	"net/http"
)

type APIError struct {
	Code  int    `json:"-"`
	Error string `json:"error"`
}

var (
	APIErrInternal = APIError{http.StatusInternalServerError, ErrInternal.Error()}
)

var (
	ErrInternal         = errors.New("internal server error")
	// ErrMethodNotAllowed = errors.New("method not allowed")
	// ErrInvalidRequest   = errors.New("invalid request")
	// ErrInvalidBody      = errors.New("invalid body")
)
