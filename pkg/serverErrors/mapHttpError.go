package serverErrors

import (
	"errors"
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/repository"
	"github.com/TheVovchenskiy/banners/internal/usecase"
	"github.com/TheVovchenskiy/banners/pkg/validator"
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
	{ErrInvalidBody, http.StatusBadRequest},
	{validator.ErrInvalidUsername, http.StatusBadRequest},
	{validator.ErrInvalidPassword, http.StatusBadRequest},
	{validator.ErrInvalidRole, http.StatusBadRequest},
	{validator.ErrInvalidAdminKey, http.StatusBadRequest},
	{repository.ErrAccountAlreadyExists, http.StatusConflict},
	{usecase.ErrInvalidLoginData, http.StatusBadRequest},
}

func MapHTTPError(err error) APIError {
	for _, allowedErr := range AllowedErrors {
		if errors.Is(err, allowedErr.Error) {
			return APIError{
				Error: err,
				Code:  allowedErr.Code,
			}
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
