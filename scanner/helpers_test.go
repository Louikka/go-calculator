package scanner

import "testing"

func TestStringify(t *testing.T) {
	tests := []struct {
		name     string
		input    []Token
		delim    string
		expected string
	}{
		{
			input: []Token{
				NewTokenNumber(0),
			},
			expected: "0",
		},
		{
			input: []Token{
				NewTokenNumber(1),
				NewTokenOperator("+"),
				NewTokenNumber(2),
			},
			delim:    " ",
			expected: "1 + 2",
		},
		{
			input: []Token{
				NewTokenWord("FUNC"),
				NewTokenPunctuation("("),
				NewTokenNumber(12.30),
				NewTokenPunctuation(")"),
			},
			expected: "FUNC(12.3)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stringify(tt.input, tt.delim)
			if s != tt.expected {
				t.Errorf("got %s, expected %s", s, tt.expected)
			}
		})
	}
}
