package serverErrors

import (
	"errors"
)

var (
	ErrInternal           = errors.New("internal server error")
	ErrInvalidQueryParams = errors.New("invalid query params")
)
