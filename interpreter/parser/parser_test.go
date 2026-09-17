package parser

import (
	"gocalc/interpreter/scanner"
	"testing"
)

func TestIsOperatorStart(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1 + 2",
			expected: "1 2 +",
		},
		{
			input:    "1 + 2 - 3",
			expected: "1 2 + 3 -",
		},
		{
			input:    "(4 + 5) * 3 - 7",
			expected: "4 5 + 3 * 7 -",
		},
		{
			input:    "(3 + 6) * (2 - 4) + 7",
			expected: "3 6 + 2 4 - * 7 +",
		},
		{
			input:    "3 + 4 * 2 / (1 - 5) ^ 2 ^ 3",
			expected: "3 4 2 * 1 5 - 2 3 ^ ^ / +",
		},
		{
			input:    "1..2",
			expected: "1 2 ..",
		},
		{
			input:    "I = 10..20",
			expected: "I 10 20 .. =",
		},
		{
			input:    "cos(1)",
			expected: "1 COS",
		},
		{
			input:    "func(1, 2, 3)",
			expected: "1 2 3 FUNC",
		},
		{
			input:    "sin(max(2, 3) / 3 * pi)",
			expected: "2 3 MAX 3 / PI * SIN",
		},
		{
			input:    "sum(i=1..5, i)",
			expected: "I 1 5 .. = I SUM",
		},
	}

	for i, test := range tests {
		tl, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("\"%s\" (case no.%d) error => %s", test.input, i, err)
		}

		rpn, err := converExpressionToRPN(tl)
		if err != nil {
			t.Errorf("\"%s\" (case no.%d) error => %s", test.input, i, err)
		}

		str := scanner.StringifyTokens(rpn, " ")
		if str != test.expected {
			t.Errorf("\"%s\" (case no.%d) => expected \"%s\", got \"%s\"", test.input, i, test.expected, str)
		}
	}
}
