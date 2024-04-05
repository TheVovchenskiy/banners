package queryManager

import "errors"

var (
	ErrUnsupportedParamType = errors.New("unsupported type of query parameter")
	ErrMissingRequiredParam = errors.New("required query paramemter is missing")
	ErrConvertingParam      = errors.New("error while converting parameter")
)
