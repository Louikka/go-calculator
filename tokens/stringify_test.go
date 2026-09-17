package token

import "testing"

func TestStringify(t *testing.T) {
	tests := []struct {
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
				NewTokenNumber(2),
				NewTokenNumber(3),
			},
			expected: "123",
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
	}

	for i, test := range tests {
		s := Stringify(test.input, test.delim)

		if s != test.expected {
			t.Errorf("(case no.%d) => got \"%s\", expected \"%s\"", i, s, test.expected)
		}
	}
}
