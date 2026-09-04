package scanner

import (
	"fmt"
	"gocalc/interpreter/lexemes"
	"slices"
	"strconv"
	"strings"
)

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
// fail).
func (s *Scanner) isEnd() bool {
	return s.pos >= len(s.s)
}

func (s *Scanner) peek(offset int) (byte, error) {
	newPos := s.pos + offset

	if newPos < 0 || newPos >= len(s.s) {
		return 0, ErrOutOfBounds
	}

	return s.s[newPos], nil
}

func (s *Scanner) next() (byte, error) {
	s.pos++
	return s.peek(0)
}

// Checks if input string is empty (length is 0).
func (s *Scanner) IsEmpty() bool {
	return len(s.s) == 0
}

/* */

type _PredicateFunc func(char, before, after byte, s string) (bool, error)

func (s *Scanner) readwhile(predicate _PredicateFunc) (string, error) {
	str := ""

	for !s.isEnd() {
		char, err := s.peek(0)
		if err != nil {
			return str, err
		}

		before, err := s.peek(-1)
		if err != nil {
			before = 0
		}

		after, err := s.peek(1)
		if err != nil {
			after = 0
		}

		predic, err := predicate(char, before, after, str)
		if err != nil {
			return str, err
		}

		if !predic {
			break
		}

		str += string(char)
		s.next()
	}

	return str, nil
}

func (s *Scanner) readNumber() (TokenNumber, error) {
	isFloat := false
	isScientific := false

	n_s, err := s.readwhile(func(char, before, after byte, _ string) (bool, error) {
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

	n_f, err := strconv.ParseFloat(n_s, 64)

	return TokenNumber{
		Value: n_f,
	}, err
}

func (s *Scanner) readWord() (TokenWord, error) {
	w, err := s.readwhile(func(char, _, _ byte, _ string) (bool, error) {
		return isLetter(char), nil
	})

	return TokenWord{
		Value: w,
	}, err
}

func (s *Scanner) readOperator() (TokenOperator, error) {
	o, err := s.readwhile(func(_, _, _ byte, str string) (bool, error) {
		if len(str) > lexemes.LONGEST_OPERATOR_LEN {
			return false, ErrTokenTooLong
		}

		if slices.Contains(lexemes.DEFINED_OPERATORS, str) {
			return false, nil
		}

		return true, nil
	})

	return TokenOperator{
		Value: o,
	}, err
}

func (s *Scanner) readPunctuation() (TokenPunctuation, error) {
	p, err := s.readwhile(func(_, _, _ byte, str string) (bool, error) {
		if len(str) > lexemes.LONGEST_PUCTUATION_LEN {
			return false, ErrTokenTooLong
		}

		if slices.Contains(lexemes.DEFINED_PUCTUATION, str) {
			return false, nil
		}

		return true, nil
	})

	return TokenPunctuation{
		Value: p,
	}, err
}

func (s *Scanner) scanNextToken() (Token, error) {
	_, err := s.readwhile(func(char, _, _ byte, _ string) (bool, error) {
		return isWhitespace(char), nil
	})
	if err != nil {
		return InvalidToken{}, err
	}

	if s.isEnd() {
		return InvalidToken{}, ErrEndOfInput
	}

	char, err := s.peek(0)
	if err != nil {
		return InvalidToken{}, err
	}

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

	return InvalidToken{}, fmt.Errorf("undefined character \"%c\"", char)
}

func Scan(s string) ([]Token, error) {
	output := []Token{}

	scanner := NewScanner(s)
	if scanner.IsEmpty() {
		return []Token{}, fmt.Errorf("Entered empty string.")
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
