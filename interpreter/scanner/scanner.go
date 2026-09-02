package scanner

import (
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

func _NewScanner(s string) _Scanner {
	return _Scanner{
		s:   strings.ToUpper(strings.TrimSpace(s)),
		pos: 0,
	}
}

func (r _Scanner) peek(offset int) (byte, error) {
	newPos := r.pos + offset

	if newPos < 0 || newPos >= len(r.s) {
		return 0, ErrOutOfBounds
	}

	return r.s[newPos], nil
}

func (r *_Scanner) next() (byte, error) {
	r.pos++

	if r.pos >= len(r.s) {
		return 0, ErrOutOfBounds
	}

	return r.s[r.pos], nil
}

// Reports if character on current position is last.
func (r _Scanner) isLast() bool {
	return r.pos == len(r.s)-1
}

// Reports if no available characters left (meaning peek() and next() will
// fail).
func (r _Scanner) isEnd() bool {
	return r.pos >= len(r.s)
}

// Checks if input string is empty (length is 0).
func (r _Scanner) IsEmpty() bool {
	return len(r.s) <= 0
}

/* */

type _PredicateFunc func(char, before, after byte, readString string) (bool, error)

func (r *_Scanner) readwhile(predicate _PredicateFunc) (string, error) {
	s := ""

	for !r.isEnd() {
		currChar, err := r.peek(0)
		if err != nil {
			return s, err
		}

		beforeChar, err := r.peek(-1)
		if err != nil {
			beforeChar = 0
		}

		afterChar, err := r.peek(1)
		if err != nil {
			afterChar = 0
		}

		predic, err := predicate(currChar, beforeChar, afterChar, s)
		if err != nil {
			return s, err
		}

		if !predic {
			break
		}

		s += string(currChar)
		r.next()
	}

	return s, nil
}

func (r *_Scanner) readNumber() (Token, error) {
	isFloat := false
	isScientific := false

	n_str, err := r.readwhile(func(char, before, after byte, _ string) (bool, error) {
		if char == '.' {
			if isFloat {
				return false, nil
			}

			isFloat = true
			return true, nil
		}

		if char == 'E' && (after == '-' || unicode.IsDigit(rune(after))) {
			if isScientific {
				return false, nil
			}

			isScientific = true
			return true, nil
		}

		if char == '-' && isScientific && before == 'E' {
			return true, nil
		}

		return unicode.IsDigit(rune(char)), nil
	})
	if err != nil {
		return Token{}, err
	}

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

	char_s := string(char)

	r.next()

	if slices.Contains(ALLOWED_OPERATORS, char_s) {
		return Token{
			Type:  TOKEN_TYPE_OPERATOR,
			Value: char_s,
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined operator \"%s\".", char_s)
	}
}

func (r *_Scanner) readVariable() (Token, error) {
	isVar := false

	v, err := r.readwhile(func(char, before, after byte, readString string) (bool, error) {
		if char == '$' {
			if unicode.IsDigit(rune(after)) {
				return false, fmt.Errorf("variables cannot start with a digit")
			}

			isVar = true
			return true, nil
		}

		if isVar && (IsLetter(char) || unicode.IsDigit(rune(char))) {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return Token{}, err
	}

	return Token{
		Type:  TOKEN_TYPE_VARIABLE,
		Value: v,
	}, nil
}

func (r *_Scanner) readKeyword() (Token, error) {
	keyw, err := r.readwhile(func(char, before, after byte, readString string) (bool, error) {
		return IsLetter(char), nil
	})
	if err != nil {
		return Token{}, err
	}

	nextT, err := r.peek(0)
	if err != nil {
		if err == ErrOutOfBounds {
			// out of bounds error => no more characters afterwards
			// therefore it is a constant
			return Token{
				Type:  TOKEN_TYPE_CONSTANT,
				Value: keyw,
			}, nil
		} else {
			return Token{}, err
		}
	}

	if string(nextT) == PUNCTUATION_LPAREN {
		// if next character is opening parenthesis, then it is a
		// function
		return Token{
			Type:  TOKEN_TYPE_FUNCTION,
			Value: keyw,
		}, nil
	} else {
		// a constant otherwise
		return Token{
			Type:  TOKEN_TYPE_CONSTANT,
			Value: keyw,
		}, nil
	}
}

func (r *_Scanner) readPunctuation() (Token, error) {
	char, err := r.peek(0)
	if err != nil {
		return Token{}, err
	}

	r.next()

	if slices.Contains(ALLOWED_PUCTUATION, string(char)) {
		return Token{
			Type:  TOKEN_TYPE_PUNCTUATION,
			Value: string(char),
		}, nil
	} else {
		return Token{}, fmt.Errorf("Undefined punctuation \"%s\".", string(char))
	}
}

func (r *_Scanner) scanNextToken() (Token, error) {
	_, err := r.readwhile(func(char, before, after byte, readString string) (bool, error) {
		return unicode.IsSpace(rune(char)), nil
	})
	if err != nil {
		return Token{}, err
	}

	if r.isEnd() {
		return Token{}, fmt.Errorf("The end of string is encountered.")
	}

	char, err := r.peek(0)
	if err != nil {
		return Token{}, err
	}

	if unicode.IsDigit(rune(char)) {
		return r.readNumber()
	}

	if slices.Contains(ALLOWED_OPERATORS, string(char)) {
		return r.readOperator()
	}

	if char == '$' {
		return r.readVariable()
	}

	if IsLetter(char) {
		return r.readKeyword()
	}

	if slices.Contains(ALLOWED_PUCTUATION, string(char)) {
		return r.readPunctuation()
	}

	return Token{}, fmt.Errorf("Undefined character \"%c\".", char)
}

func Scan(s string) ([]Token, error) {
	output := []Token{}

	scanner := _NewScanner(s)
	if scanner.IsEmpty() {
		return []Token{}, fmt.Errorf("Entered empty string.")
	}

	for !scanner.isEnd() {
		t, err := scanner.scanNextToken()
		if err != nil {
			return []Token{}, err
		}

		output = append(output, t)
	}

	return output, nil
}
