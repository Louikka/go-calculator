package parser

import "errors"

var (
	ErrEmptyExpr = errors.New("empty expression")

	ErrMismatchedParenthesis = errors.New("mismatched parenthesis")

	ErrNotAnUnary = errors.New("not an unary expression")
	ErrNotABinary = errors.New("not a binary expression")
)
