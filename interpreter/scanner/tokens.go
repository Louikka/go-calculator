package scanner

import (
	"gocalc/interpreter/lexemes"
	"math"
	"strconv"
	"strings"
)

type Token interface {
	// String representation of a token type.
	Type() string
	// String representation of a token value.
	ToString() string
}

// invalid

type InvalidToken struct {
	//
}

func (t InvalidToken) Type() string {
	return "INVALID"
}

func (t InvalidToken) ToString() string {
	return ""
}

// number

type TokenNumber struct {
	Value float64
}

func NewTokenNumber(v float64) TokenNumber {
	return TokenNumber{
		Value: v,
	}
}

func (t TokenNumber) Type() string {
	return "NUMBER"
}

func (t TokenNumber) ToString() string {
	return strconv.FormatFloat(t.Value, 'f', -1, 64)
}

func (t TokenNumber) IsInt() bool {
	return t.Value == math.Trunc(t.Value)
}

// word

const (
	WORD_KIND_UNDEFINED = ""
	WORD_KIND_CONSTANT  = "CONSTANT"
	WORD_KIND_VARIABLE  = "VARIABLE"
	WORD_KIND_FUNCTION  = "FUNCTION"
)

// Represents constants, variables and functions (basically, everything that
// starts with a letter).
type TokenWord struct {
	Value string
	Kind  string
}

func (t TokenWord) Type() string {
	if t.Kind == WORD_KIND_UNDEFINED {
		return "WORD"
	} else {
		return t.Kind
	}
}

func (t TokenWord) ToString() string {
	return t.Value
}

// operator

type TokenOperator struct {
	Value string
}

func NewTokenOperator(v string) TokenOperator {
	return TokenOperator{
		Value: v,
	}
}

func (t TokenOperator) Type() string {
	return "OPERATOR"
}

func (t TokenOperator) ToString() string {
	return t.Value
}

func (t TokenOperator) Definition() lexemes.OperatorDefinition {
	def, _ := lexemes.IsOperator(t.Value)
	return def
}

// punctuation

type TokenPunctuation struct {
	Value string
}

func NewTokenPunctuation(v string) TokenPunctuation {
	return TokenPunctuation{
		Value: v,
	}
}

func (t TokenPunctuation) Type() string {
	return "PUNCTUATION"
}

func (t TokenPunctuation) ToString() string {
	return t.Value
}

func (t TokenPunctuation) IsParenthesis() bool {
	return (t.Value == "(") || (t.Value == ")")
}

func (t TokenPunctuation) IsLeftParenthesis() bool {
	return t.Value == "("
}

/* Helpers *******************************************************************/

func StringifyTokens(tl []Token, delimeter string) string {
	var s strings.Builder

	tlLastIndex := len(tl) - 1

	for i, t := range tl {
		a := t.ToString()

		if i < tlLastIndex {
			a += delimeter
		}

		s.WriteString(a)
	}

	return s.String()
}
