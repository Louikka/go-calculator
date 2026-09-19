package scanner

import (
	"math"
	"strconv"
)

type Token interface {
	// String representation of a token type.
	Type() string
	// String representation of a token value.
	ToString() string
}

// Invalid token //----------------------------------------------------------//

type TokenInvalid struct {
	//
}

func (t TokenInvalid) Type() string {
	return "INVALID"
}

func (t TokenInvalid) ToString() string {
	return ""
}

// Number token //-----------------------------------------------------------//

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

// Word token //-------------------------------------------------------------//

const (
	WORD_KIND_UNSPECIFIED = ""
	WORD_KIND_IDENTIFIER  = "IDENTIFIER"
	WORD_KIND_FUNCTION    = "FUNCTION"
)

// Represents constants, variables and functions (basically, everything that
// starts with a letter).
type TokenWord struct {
	Value string
	Kind  string
}

func NewTokenWord(v string) TokenWord {
	return TokenWord{
		Value: v,
	}
}

func (t TokenWord) Type() string {
	if t.Kind == WORD_KIND_UNSPECIFIED {
		return "WORD"
	} else {
		return t.Kind
	}
}

func (t TokenWord) ToString() string {
	return t.Value
}

// Operator token //---------------------------------------------------------//

type TokenOperator struct {
	Value         string
	Precedence    int
	Associativity string
}

func NewTokenOperator(v string) TokenOperator {
	oper, isOper := IsOperator(v)
	if isOper {
		return TokenOperator{
			Value:         oper.Value,
			Precedence:    oper.Precedence,
			Associativity: oper.Associativity,
		}
	} else {
		return TokenOperator{
			Value: v,
		}
	}
}

func (t TokenOperator) Type() string {
	return "OPERATOR"
}

func (t TokenOperator) ToString() string {
	return t.Value
}

// Punctuation token //------------------------------------------------------//

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

func (t TokenPunctuation) IsLeftParenthesis() bool {
	return t.Value == "("
}
