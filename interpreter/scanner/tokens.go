package scanner

type TokenType string

const (
	TOKEN_TYPE_NUMBER      TokenType = "NUMBER"
	TOKEN_TYPE_OPERATOR    TokenType = "OPERATOR"
	TOKEN_TYPE_CONSTANT    TokenType = "CONSTANT"
	TOKEN_TYPE_FUNCTION    TokenType = "FUNCTION"
	TOKEN_TYPE_PUNCTUATION TokenType = "PUNCTUATION"
)

type Token struct {
	Type  TokenType
	Value any
}

var ALLOWED_OPERATORS = []byte{'+', '-', '*', '/', '^'}

var ALLOWED_CONSTANTS = []string{"PI", "E", "PHI"}

var ALLOWED_FUNCTIONS = []string{"SIN", "COS", "TAN", "ATAN", "EXP", "ABS", "LOG", "LN", "SQRT"}

var ALLOWED_PUCTUATION = []byte{'(', ')'}
