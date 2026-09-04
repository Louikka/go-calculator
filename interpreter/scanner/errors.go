package scanner

import "errors"

var (
	ErrOutOfBounds  = errors.New("position is out of bounds")
	ErrEndOfInput   = errors.New("end of input encountered")
	ErrTokenTooLong = errors.New("token length is too long")
)
