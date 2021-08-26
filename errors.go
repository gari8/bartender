package bartender

import "errors"

var (
	// ErrEmptyField struct has no field
	ErrEmptyField = errors.New("struct has no field")
	// ErrDifferentType type is different
	ErrDifferentType = errors.New("type is different")
)
