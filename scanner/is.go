package scanner

import (
	"unicode"
)

// Reports whether the byte b is a digit.
func isDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}

// Reports whether the byte b is a latin letter.
func isLetter(b byte) bool {
	return unicode.IsLetter(rune(b)) && unicode.Is(unicode.Latin, rune(b))
}

// Reports whether the byte b is a whitespace.
func isWhitespace(b byte) bool {
	return unicode.IsSpace(rune(b))
}

func isWordStart(b byte) bool {
	return isLetter(b)
}

// Reports whether the byte b can be part of the word (except start).
func isWordBody(b byte) bool {
	return isLetter(b) || isDigit(b) || b == '_'
}

func isOperatorStart(b byte) bool {
	for _, oper := range DEFINED_OPERATORS {
		if oper.Value[0] == b {
			return true
		}
	}

	return false
}

func isPunctuationStart(b byte) bool {
	for _, punc := range DEFINED_PUCTUATION {
		if punc.Value[0] == b {
			return true
		}
	}

	return false
}

func isLeftParenthesis(b byte) bool {
	return string(b) == "("
}
