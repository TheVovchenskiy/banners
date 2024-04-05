package serverErrors

import ()

var APIErrors = map[error]APIError{
	// ErrMethodNotAllowed: http.StatusMethodNotAllowed,
	// ErrInvalidRequest:   http.StatusBadRequest,
	// ErrInvalidBody:      http.StatusBadRequest,
}

func MapHTTPError(err error) (apiError APIError) {
	if err == nil {
		err = ErrInternal
	}

	apiError, ok := APIErrors[err]
	if !ok {
		apiError = APIErrInternal
	}

	return
}
