package serverErrors

import (
	"errors"
	"net/http"
)

type APIError struct {
	Code  int   `json:"-"`
	Error error `json:"error"`
}

var (
	APIErrInternal = APIError{http.StatusInternalServerError, ErrInternal}
)

var (
	ErrInternal = errors.New("internal server error")
)
