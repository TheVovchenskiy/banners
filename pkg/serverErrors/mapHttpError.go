package serverErrors

import ()

var APIErrors = map[error]APIError{
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
