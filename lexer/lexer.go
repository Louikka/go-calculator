package lexer

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/* Tokens ********************************************************************/

type Token struct {
	Type  string
	Value any
}

// 1  2  3.14  -45.01  0.12E4 5.886E-2  etc..
type NumberToken struct {
	Type  string
	Value float64
}

var ALLOWED_OPERATORS = []rune{'+', '-', '*', '/', '^'}

type OperatorToken struct {
	Type  string
	Value rune
}

var ALLOWED_KEYWORDS = []string{"POW", "SQRT", "ABS"}

type KeywordToken struct {
	Type  string
	Value string
}

var ALLOWED_PUCTUATION = []rune{'(', ')'}

type PunctuationToken struct {
	Type  string
	Value string
}

/* Custom string reader ******************************************************/

type StringReader struct {
	s   string
	pos uint
}

func (r *StringReader) Peek(offset int) (byte, error) {
	newPos := int(r.pos) + offset

	if newPos < 0 || newPos >= len(r.s) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[newPos], nil
}

func (r *StringReader) Next() (byte, error) {
	r.pos++
	if r.pos >= uint(len(r.s)) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[r.pos], nil
}

func (r *StringReader) IsEndOfString() bool {
	return r.pos >= uint(len(r.s))
}

/* Lexer main ****************************************************************/

type __PredicateFunc func(byte, byte, byte, string) bool

func readWhile(r *StringReader, predicate __PredicateFunc) string {
	s := ""

	for !r.IsEndOfString() {
		currChar, err := r.Peek(0)
		if err != nil {
			fmt.Println(err)
			break
		}
		beforeChar, err := r.Peek(-1)
		if err != nil {
			beforeChar = 0
		}
		afterChar, err := r.Peek(1)
		if err != nil {
			afterChar = 0
		}
		if !predicate(currChar, beforeChar, afterChar, s) {
			break
		}

		s += string(currChar)
		r.Next()
	}

	return s
}

func readNumber(r *StringReader) (Token, error) {
	isFloat := false
	isScientific := false

	n_asStr := readWhile(r, func(char byte, before byte, after byte, readString string) bool {
		if char == '.' {
			if isFloat {
				return false
			}

			isFloat = true
			return true
		}

		if char == 'E' && (after == '-' || unicode.IsDigit(rune(after))) {
			if isScientific {
				return false
			}

			isScientific = true
			return true
		}

		if char == '-' && isScientific && before == 'E' {
			return true
		}

		return unicode.IsDigit(rune(char))
	})

	n_asFloat, err := strconv.ParseFloat(n_asStr, 64)
	if err != nil {
		return Token{}, err
	}

	return Token{
		Type:  "NUMBER",
		Value: n_asFloat,
	}, nil
}

func readOperator(r *StringReader) (Token, error) {
	char, err := r.Peek(0)
	if err != nil {
		return Token{}, err
	}

	r.Next()

	if slices.Contains(ALLOWED_OPERATORS, rune(char)) {
		return Token{
			Type:  "OPERATOR",
			Value: char,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined operator \"%s\".", string(char))
	}
}

func readKeyword(r *StringReader) (Token, error) {
	keyw := readWhile(r, func(char byte, before byte, after byte, readString string) bool {
		return unicode.IsLetter(rune(char)) && unicode.Is(unicode.Latin, rune(char))
	})

	if slices.Contains(ALLOWED_KEYWORDS, keyw) {
		return Token{
			Type:  "KEYWORD",
			Value: keyw,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined keyword \"%s\".", string(keyw))
	}
}

func readPunctuation(r *StringReader) (Token, error) {
	char, err := r.Peek(0)
	if err != nil {
		return Token{}, err
	}

	r.Next()

	if slices.Contains(ALLOWED_PUCTUATION, rune(char)) {
		return Token{
			Type:  "PUCTUATION",
			Value: char,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined punctuation \"%s\".", string(char))
	}
}

func analyzeNextToken(r *StringReader) (Token, error) {
	readWhile(r, func(char byte, before byte, after byte, readString string) bool {
		return unicode.IsSpace(rune(char))
	})

	if r.IsEndOfString() {
		return Token{}, fmt.Errorf("The end of string is encountered.")
	}

	char, err := r.Peek(0)
	if err != nil {
		return Token{}, err
	}

	if unicode.IsDigit(rune(char)) {
		return readNumber(r)
	}

	if slices.Contains(ALLOWED_OPERATORS, rune(char)) {
		return readOperator(r)
	}

	if unicode.IsLetter(rune(char)) && unicode.Is(unicode.Latin, rune(char)) {
		return readKeyword(r)
	}

	if slices.Contains(ALLOWED_PUCTUATION, rune(char)) {
		return readPunctuation(r)
	}

	return Token{}, fmt.Errorf("Undefined character \"%s\".", string(char))
}

func Analyse(s string) ([]any, error) {
	output := []any{}

	s_prep := strings.ToUpper(strings.TrimSpace(s))
	reader := StringReader{
		s:   s_prep,
		pos: 0,
	}

	for !reader.IsEndOfString() {
		t, err := analyzeNextToken(&reader)
		if err != nil {
			return []any{}, err
		}

		output = append(output, t)
	}

	return output, nil
}
