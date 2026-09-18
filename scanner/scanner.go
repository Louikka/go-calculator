package scanner

import (
	"fmt"
	"strconv"
	"strings"
)

//---------------------------------------------------------------------------//

type Scanner struct {
	s   string
	pos int
}

func NewScanner(s string) Scanner {
	return Scanner{
		s:   strings.ToUpper(strings.TrimSpace(s)),
		pos: 0,
	}
}

// Reports if character on current position is last.
func (s *Scanner) isLast() bool {
	return s.pos == len(s.s)-1
}

// Reports if no available characters left (meaning peek() and next() will
// panic).
func (s *Scanner) isEnd() bool {
	return s.pos >= len(s.s)
}

func (s *Scanner) peek(offset int) byte {
	newPos := s.pos + offset

	if newPos < 0 || newPos >= len(s.s) {
		panic("scanner error: index out of range")
	}

	return s.s[newPos]
}

func (s *Scanner) next() byte {
	s.pos++
	return s.peek(0)
}

// Checks if input string is empty (length is 0).
func (s *Scanner) IsEmpty() bool {
	return len(s.s) == 0
}

//---------------------------------------------------------------------------//

type _PredicateFunc func(char, before, after byte, s string) (bool, error)

func (s *Scanner) readwhile(predicate _PredicateFunc) (string, error) {
	str := ""

	for !s.isEnd() {
		var char, before, after byte = 0, 0, 0

		char = s.peek(0)
		if s.pos > 0 {
			before = s.peek(-1)
		}
		if !s.isLast() {
			after = s.peek(1)
		}

		predic, err := predicate(char, before, after, str)
		if err != nil {
			return str, err
		}

		if !predic {
			break
		}

		str += string(char)

		s.pos++
	}

	return str, nil
}

// Skips whitespace characters.
func (s *Scanner) skipwsp() error {
	_, err := s.readwhile(func(char, _, _ byte, _ string) (bool, error) {
		return isWhitespace(char), nil
	})

	return err
}

//---------------------------------------------------------------------------//

func (s *Scanner) readNumber() (TokenNumber, error) {
	isFloat := false
	isScientific := false

	r, err := s.readwhile(func(char, before, after byte, _ string) (bool, error) {
		if char == '.' {
			if isFloat {
				return false, nil
			}

			if isDigit(after) {
				isFloat = true
				return true, nil
			} else {
				return false, nil
			}
		}

		if char == 'E' && (after == '-' || isDigit(after)) {
			if isScientific {
				return false, nil
			}

			isScientific = true
			return true, nil
		}

		if char == '-' && isScientific && before == 'E' {
			return true, nil
		}

		return isDigit(char), nil
	})
	if err != nil {
		return TokenNumber{}, err
	}

	n, err := strconv.ParseFloat(r, 64)

	return NewTokenNumber(n), err
}

func (s *Scanner) readWord() (TokenWord, error) {
	t := TokenWord{}

	w, err := s.readwhile(func(char, _, _ byte, s string) (bool, error) {
		return isLetter(char) || (isDigit(char) && len(s) > 0), nil
	})
	if err != nil {
		return t, err
	}

	t.Value = w

	err = s.skipwsp()
	if err != nil {
		return t, err
	}

	if !s.isEnd() && isLeftParenthesis(s.peek(0)) {
		t.Kind = WORD_KIND_FUNCTION
	}

	return t, err
}

func (s *Scanner) readOperator() (TokenOperator, error) {
	o, err := s.readwhile(func(_, _, _ byte, str string) (bool, error) {
		if len(str) > LongestOperatorLen {
			return false, ErrTokenTooLong
		}

		if _, isOper := IsOperator(str); isOper {
			return false, nil
		}

		return true, nil
	})

	return NewTokenOperator(o), err
}

func (s *Scanner) readPunctuation() (TokenPunctuation, error) {
	p, err := s.readwhile(func(_, _, _ byte, str string) (bool, error) {
		if len(str) > LongestPunctuationLen {
			return false, ErrTokenTooLong
		}

		if _, isPunc := IsPunctuation(str); isPunc {
			return false, nil
		}

		return true, nil
	})

	return TokenPunctuation{
		Value: p,
	}, err
}

func (s *Scanner) scanNextToken() (Token, error) {
	err := s.skipwsp()
	if err != nil {
		return TokenInvalid{}, err
	}

	if s.isEnd() {
		return TokenInvalid{}, ErrEndOfInput
	}

	char := s.peek(0)

	if isDigit(char) {
		return s.readNumber()
	}

	if isLetter(char) {
		return s.readWord()
	}

	if isOperatorStart(char) {
		return s.readOperator()
	}

	if isPunctuationStart(char) {
		return s.readPunctuation()
	}

	return TokenInvalid{}, fmt.Errorf("undefined character \"%c\"", char)
}

//---------------------------------------------------------------------------//

func Scan(s string) ([]Token, error) {
	output := []Token{}

	scanner := NewScanner(s)
	if scanner.IsEmpty() {
		return []Token{}, nil
	}

	for !scanner.isEnd() {
		t, err := scanner.scanNextToken()
		if err != nil {
			return output, err
		}

		output = append(output, t)
	}

	return output, nil
}
