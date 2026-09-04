package parser

import "errors"

var (
	ErrMismatchedParenthesis = errors.New("mismatched parenthesis")
	ErrVarAsConst            = errors.New("expected variable, but got constant")
)
