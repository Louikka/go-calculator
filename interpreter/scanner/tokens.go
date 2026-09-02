package scanner

type TokenType string

const (
	TOKEN_TYPE_NUMBER      TokenType = "NUMBER"
	TOKEN_TYPE_OPERATOR    TokenType = "OPERATOR"
	TOKEN_TYPE_VARIABLE    TokenType = "VARIABLE"
	TOKEN_TYPE_CONSTANT    TokenType = "CONSTANT"
	TOKEN_TYPE_FUNCTION    TokenType = "FUNCTION"
	TOKEN_TYPE_PUNCTUATION TokenType = "PUNCTUATION"
)

type Token struct {
	Type  TokenType
	Value any
}

func NewToken(t TokenType, v any) Token {
	return Token{
		Type:  t,
		Value: v,
	}
}

var ALLOWED_OPERATORS = []string{"+", "-", "*", "/", "^"}

const (
	OPERATOR_ADD = "+"
	OPERATOR_SUB = "-"
	OPERATOR_MUL = "*"
	OPERATOR_DIV = "/"
	OPERATOR_POW = "^"
)

var ALLOWED_CONSTANTS = []string{"PI", "E", "PHI"}

const (
	CONSTANT_PI  = "PI"
	CONSTANT_E   = "E"
	CONSTANT_PHI = "PHI"
)

var ALLOWED_FUNCTIONS = []string{"SIN", "COS", "TAN", "ATAN", "EXP", "ABS", "LOG", "LN", "SQRT"}

const (
	FUNCTION_SIN  = "SIN"
	FUNCTION_COS  = "COS"
	FUNCTION_TAN  = "TAN"
	FUNCTION_ATAN = "ATAN"
	FUNCTION_EXP  = "EXP"
	FUNCTION_ABS  = "ABS"
	FUNCTION_LOG  = "LOG"
	FUNCTION_LN   = "LN"
	FUNCTION_SQRT = "SQRT"
)

var ALLOWED_PUCTUATION = []string{"(", ")"}

const (
	PUNCTUATION_LPAREN = "("
	PUNCTUATION_RPAREN = ")"
)

/* Helper functions **********************************************************/

func NewNumberToken(v float64) Token {
	return Token{
		Type:  TOKEN_TYPE_NUMBER,
		Value: v,
	}
}

func NewOperatorToken(v string) Token {
	return Token{
		Type:  TOKEN_TYPE_OPERATOR,
		Value: v,
	}
}

func NewPunctuationToken(v string) Token {
	return Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: v,
	}
}
