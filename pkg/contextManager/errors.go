package contextManager

import "errors"

var (
	ErrNoValueInContext = errors.New("no key: value in context")
)
