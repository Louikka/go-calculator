package scanner

import (
	"gocalc/interpreter/lexemes"
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

func isOperatorStart(b byte) bool {
	for _, s := range lexemes.DEFINED_OPERATORS {
		if s[0] == b {
			return true
		}
	}

	return false
}

func isPunctuationStart(b byte) bool {
	for _, s := range lexemes.DEFINED_PUCTUATION {
		if s[0] == b {
			return true
		}
	}

	return false
}
