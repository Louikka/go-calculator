package lexer

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/* Types and constants *******************************************************/

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

/*
type NumberToken struct {
	Type  TokenType
	Value float64
}
type OperatorToken struct {
	Type  TokenType
	Value rune
}
type KeywordToken struct {
	Type  TokenType
	Value string
}
type PunctuationToken struct {
	Type  TokenType
	Value rune
}
*/

var ALLOWED_OPERATORS = []byte{'+', '-', '*', '/', '^'}

var ALLOWED_CONSTANTS = []string{"PI", "E" /* "TAU", "PHI", */}

var ALLOWED_FUNCTIONS = []string{"SIN", "COS", "TAN", "ATAN", "EXP", "ABS", "LOG", "LN", "SQRT"}

var ALLOWED_PUCTUATION = []byte{'(', ')'}

/* Custom string reader ******************************************************/

type _StringReader struct {
	s   string
	pos uint
}

func (r *_StringReader) peek(offset int) (byte, error) {
	newPos := int(r.pos) + offset

	if newPos < 0 || newPos >= len(r.s) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[newPos], nil
}

func (r *_StringReader) next() (byte, error) {
	r.pos++

	if r.pos >= uint(len(r.s)) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[r.pos], nil
}

func (r *_StringReader) isEndOfString() bool {
	return r.pos >= uint(len(r.s))
}

/* Helpers *******************************************************************/

func StringifyTokens(tl []Token) string {
	s := ""

	for _, v := range tl {
		switch v.Type {
		case TOKEN_TYPE_NUMBER:
			n := v.Value.(float64)
			s += strconv.FormatFloat(n, 'f', -1, 64)
		case TOKEN_TYPE_OPERATOR, TOKEN_TYPE_PUNCTUATION:
			s += string(v.Value.(byte)) //fmt.Sprintf("%s", v.Value)
		case TOKEN_TYPE_CONSTANT, TOKEN_TYPE_FUNCTION:
			s += v.Value.(string)
		}
	}

	return s
}

/* Lexer main ****************************************************************/

type __PredicateFunc func(byte, byte, byte, string) bool

func readWhile(r *_StringReader, predicate __PredicateFunc) string {
	s := ""

	for !r.isEndOfString() {
		currChar, err := r.peek(0)
		if err != nil {
			fmt.Println(err)
			break
		}
		beforeChar, err := r.peek(-1)
		if err != nil {
			beforeChar = 0
		}
		afterChar, err := r.peek(1)
		if err != nil {
			afterChar = 0
		}
		if !predicate(currChar, beforeChar, afterChar, s) {
			break
		}

		s += string(currChar)
		r.next()
	}

	return s
}

func readNumber(r *_StringReader) (Token, error) {
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
		Type:  TOKEN_TYPE_NUMBER,
		Value: n_asFloat,
	}, nil
}

func readOperator(r *_StringReader) (Token, error) {
	char, err := r.peek(0)
	if err != nil {
		return Token{}, err
	}

	r.next()

	if slices.Contains(ALLOWED_OPERATORS, char) {
		return Token{
			Type:  TOKEN_TYPE_OPERATOR,
			Value: char,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined operator \"%s\".", string(char))
	}
}

func readKeyword(r *_StringReader) (Token, error) {
	keyw := readWhile(r, func(char byte, before byte, after byte, readString string) bool {
		return unicode.IsLetter(rune(char)) && unicode.Is(unicode.Latin, rune(char))
	})

	if slices.Contains(ALLOWED_CONSTANTS, keyw) {
		return Token{
			Type:  TOKEN_TYPE_CONSTANT,
			Value: keyw,
		}, nil
	} else if slices.Contains(ALLOWED_FUNCTIONS, keyw) {
		return Token{
			Type:  TOKEN_TYPE_FUNCTION,
			Value: keyw,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined keyword \"%s\".", string(keyw))
	}
}

func readPunctuation(r *_StringReader) (Token, error) {
	char, err := r.peek(0)
	if err != nil {
		return Token{}, err
	}

	r.next()

	if slices.Contains(ALLOWED_PUCTUATION, char) {
		return Token{
			Type:  TOKEN_TYPE_PUNCTUATION,
			Value: char,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined punctuation \"%s\".", string(char))
	}
}

func analyzeNextToken(r *_StringReader) (Token, error) {
	readWhile(r, func(char byte, before byte, after byte, readString string) bool {
		return unicode.IsSpace(rune(char))
	})

	if r.isEndOfString() {
		return Token{}, fmt.Errorf("The end of string is encountered.")
	}

	char, err := r.peek(0)
	if err != nil {
		return Token{}, err
	}

	if unicode.IsDigit(rune(char)) {
		return readNumber(r)
	}

	if slices.Contains(ALLOWED_OPERATORS, char) {
		return readOperator(r)
	}

	if unicode.IsLetter(rune(char)) && unicode.Is(unicode.Latin, rune(char)) {
		return readKeyword(r)
	}

	if slices.Contains(ALLOWED_PUCTUATION, char) {
		return readPunctuation(r)
	}

	return Token{}, fmt.Errorf("Undefined character \"%s\".", string(char))
}

func Analyse(s string) ([]Token, error) {
	output := []Token{}

	s_prep := strings.ToUpper(strings.TrimSpace(s))
	reader := _StringReader{
		s:   s_prep,
		pos: 0,
	}

	for !reader.isEndOfString() {
		t, err := analyzeNextToken(&reader)
		if err != nil {
			return []Token{}, err
		}

		output = append(output, t)
	}

	return output, nil
}
