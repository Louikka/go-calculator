package parser

import "errors"

var (
	ErrMismatchedParenthesis = errors.New("mismatched parenthesis")

	ErrNotABinary = errors.New("not a binary expression")
)
