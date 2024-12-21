package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("Expression is not valid")
	ErrInternalServer = errors.New("Internal server error")
	ErrMethodNotAllowed = errors.New("Method not allowed")
)
