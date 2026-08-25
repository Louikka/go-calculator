package scanner

import (
	"strconv"
	"strings"
)

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

/* Helpers *******************************************************************/

func StringifyTokens(tl []Token) string {
	var s strings.Builder

	for _, v := range tl {
		switch v.Type {
		case TOKEN_TYPE_NUMBER:
			n := v.Value.(float64)
			s.WriteString(strconv.FormatFloat(n, 'f', -1, 64))
		case TOKEN_TYPE_OPERATOR, TOKEN_TYPE_PUNCTUATION:
			s.WriteString(string(v.Value.(byte))) //fmt.Sprintf("%s", v.Value)
		case TOKEN_TYPE_CONSTANT, TOKEN_TYPE_FUNCTION:
			s.WriteString(v.Value.(string))
		}
	}

	return s.String()
}
