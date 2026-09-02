package scanner

import (
	"math"
	"strconv"
)

type Token interface {
	Type() string
	ToString() string
}

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

func (t TokenPunctuation) Type() string {
	return "PUNCTUATION"
}

func (t TokenPunctuation) ToString() string {
	return t.Value
}
