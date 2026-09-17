package scanner

import "errors"

var (
	ErrEndOfInput   = errors.New("end of input encountered")
	ErrTokenTooLong = errors.New("token length is too long")
)
