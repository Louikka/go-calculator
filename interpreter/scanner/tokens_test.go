package scanner

import (
	"testing"
)

func TestStringifyTokens(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "1 + 2",
			expected: "1+2",
		},
		{
			input:    "(1 + 2) * 3",
			expected: "(1+2)*3",
		},
		{
			input:    "(1 + 2) * 3 / 4",
			expected: "(1+2)*3/4",
		},
		{
			input:    "(1 + (2 * 3) / 4)^5",
			expected: "(1+(2*3)/4)^5",
		},
		{
			input:    "SIN(1)",
			expected: "SIN(1)",
		},
		{
			input:    "COS( 1 + 2 )",
			expected: "COS(1+2)",
		},
		{
			input:    "PI * 4 - 5.12",
			expected: "PI*4-5.12",
		},
		{
			input:    "1..2",
			expected: "1..2",
		},
		{
			input:    "SUM(I=1..2, I * 2)",
			expected: "SUM(I=1..2,I*2)",
		},
		{
			input:    "1 2.3 4e5 6e-7 + - * / ^ = .. PI E PHI SIN COS TAN SQRT ( ) ,",
			expected: "12.34000000.0000006+-*/^=..PIEPHISINCOSTANSQRT(),",
		},
	}

	for i, test := range tests {
		tl, err := Scan(test.input)
		if err != nil {
			t.Errorf("(case no.%d) => Scan error : %s", i, err)
		}

		s := StringifyTokens(tl)
		if s != test.expected {
			t.Errorf("(case no.%d) => got %s, expected %s", i, s, test.expected)
		}
	}
}
