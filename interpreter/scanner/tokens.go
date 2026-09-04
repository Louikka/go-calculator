package scanner

import (
	"math"
	"strconv"
	"strings"
)

type Token interface {
	Type() string
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

type TokenWord struct {
	Value string
}

func (t TokenWord) Type() string {
	return "WORD"
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

/* Helpers *******************************************************************/

func StringifyTokens(tl []Token) string {
	var s strings.Builder

	for _, t := range tl {
		s.WriteString(t.ToString())
	}

	return s.String()
}
