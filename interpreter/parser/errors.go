package parser

import "errors"

var (
	ErrEmptyExpression       = errors.New("empty expression")
	ErrNotABinaryExpression  = errors.New("not a binary expression")
	ErrMismatchedParenthesis = errors.New("mismatched parenthesis")
	ErrVarAsConst            = errors.New("constant used as variable")
)
