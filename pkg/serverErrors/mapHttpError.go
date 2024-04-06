package serverErrors

import (
	"errors"
	"net/http"
)

type APIError struct {
	Error error
	Code  int
}

// var APIErrorStatusCodes = map[error]int{
// 	ErrInternal:           http.StatusInternalServerError,
// 	ErrInvalidQueryParams: http.StatusBadRequest,
// }

var AllowedErrors []APIError = []APIError{
	{ErrInternal, http.StatusInternalServerError},
	{ErrInvalidQueryParams, http.StatusBadRequest},
}

func MapHTTPError(err error) APIError {
	for _, allowedErr := range AllowedErrors {
		if errors.Is(err, allowedErr.Error) {
			return allowedErr
		}
	}

	return APIError{
		Error: ErrInternal,
		Code:  http.StatusInternalServerError,
	}
	// status, ok := APIErrorStatusCodes[err]
	// if !ok {
	// 	apiError = APIError{
	// 		Error: ErrInternal,
	// 		Code:  http.StatusInternalServerError,
	// 	}
	// } else {
	// 	apiError = APIError{
	// 		Code:  status,
	// 		Error: err,
	// 	}
	// }

	// return
}
