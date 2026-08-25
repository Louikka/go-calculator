package scanner

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type _Scanner struct {
	s   string
	pos int
}

func _NewScanner(s string) *_Scanner {
	return &_Scanner{
		s:   s,
		pos: 0,
	}
}

func (r *_Scanner) peek(offset int) (byte, error) {
	newPos := int(r.pos) + offset

	if newPos < 0 || newPos >= len(r.s) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[newPos], nil
}

func (r *_Scanner) next() (byte, error) {
	r.pos++

	if r.pos >= len(r.s) {
		return 0, errors.New("Position is out of bounds.")
	}

	return r.s[r.pos], nil
}

func (r *_Scanner) isEndOfString() bool {
	return r.pos >= len(r.s)
}

type _PredicateFunc func(byte, byte, byte, string) bool

func (r *_Scanner) readwhile(predicate _PredicateFunc) string {
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

func (r *_Scanner) readNumber() (Token, error) {
	isFloat := false
	isScientific := false

	n_str := r.readwhile(func(char byte, before byte, after byte, _ string) bool {
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

	n_f, err := strconv.ParseFloat(n_str, 64)
	if err != nil {
		return Token{}, err
	}

	return Token{
		Type:  TOKEN_TYPE_NUMBER,
		Value: n_f,
	}, nil
}

func (r *_Scanner) readOperator() (Token, error) {
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

func (r *_Scanner) readKeyword() (Token, error) {
	keyw := r.readwhile(func(char byte, before byte, after byte, readString string) bool {
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

func (r *_Scanner) readPunctuation() (Token, error) {
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

func (r *_Scanner) scanNextToken() (Token, error) {
	r.readwhile(func(char byte, before byte, after byte, readString string) bool {
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
		return r.readNumber()
	}

	if slices.Contains(ALLOWED_OPERATORS, char) {
		return r.readOperator()
	}

	if unicode.IsLetter(rune(char)) && unicode.Is(unicode.Latin, rune(char)) {
		return r.readKeyword()
	}

	if slices.Contains(ALLOWED_PUCTUATION, char) {
		return r.readPunctuation()
	}

	return Token{}, fmt.Errorf("Undefined character \"%s\".", string(char))
}

func Scan(s string) ([]Token, error) {
	output := []Token{}

	s_prep := strings.ToUpper(strings.TrimSpace(s))
	if len(s_prep) == 0 {
		return []Token{}, fmt.Errorf("Entered empty string.")
	}

	scanner := _NewScanner(s_prep)

	for !scanner.isEndOfString() {
		t, err := scanner.scanNextToken()
		if err != nil {
			return []Token{}, err
		}

		output = append(output, t)
	}

	return output, nil
}
