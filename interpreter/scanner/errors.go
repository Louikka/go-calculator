package scanner

import "errors"

var (
	ErrOutOfBounds      = errors.New("position is out of bounds")
	ErrIllegalCharacter = errors.New("illegal character")
	ErrTokenTooLong     = errors.New("token length is too long")
)
