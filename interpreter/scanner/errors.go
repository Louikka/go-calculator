package scanner

import "errors"

var (
	ErrOutOfBounds      = errors.New("position is out of bounds")
	ErrEndOfInput       = errors.New("end of input encountered")
	ErrIllegalCharacter = errors.New("illegal character")
	ErrTokenTooLong     = errors.New("token length is too long")
)
