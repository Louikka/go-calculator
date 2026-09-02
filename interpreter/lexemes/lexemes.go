package lexemes

import "gocalc/lib"

var DEFINED_CONSTANTS = []string{"PI", "E", "PHI"}

const (
	CONSTANT_PI  = "PI"
	CONSTANT_E   = "E"
	CONSTANT_PHI = "PHI"
)

var DEFINED_FUNCTIONS = []string{"SIN", "COS", "TAN", "ATAN", "EXP", "ABS", "LOG", "LN", "SQRT", "SUM"}

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

var DEFINED_OPERATORS = []string{"+", "-", "*", "/", "^", ".."}

// Length of the longest operator (in bytes).
var LONGEST_OPERATOR_LEN = lib.LongestStringLenInSlice(DEFINED_OPERATORS)

const (
	OPERATOR_ADD   = "+"
	OPERATOR_SUB   = "-"
	OPERATOR_MUL   = "*"
	OPERATOR_DIV   = "/"
	OPERATOR_POW   = "^"
	OPERATOR_RANGE = ".."
)

var DEFINED_PUCTUATION = []string{"(", ")", ","}

// Length of the longest punctuation (in bytes).
var LONGEST_PUCTUATION_LEN = lib.LongestStringLenInSlice(DEFINED_PUCTUATION)

const (
	PUNCTUATION_LPAREN = "("
	PUNCTUATION_RPAREN = ")"
	PUNCTUATION_COMMA  = ","
)
