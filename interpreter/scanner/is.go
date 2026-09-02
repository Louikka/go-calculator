package scanner

import "unicode"

// Reports whether the byte b is a latin letter.
func IsLetter(b byte) bool {
	return unicode.IsLetter(rune(b)) && unicode.Is(unicode.Latin, rune(b))
}
