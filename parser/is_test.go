package parser

import (
	"gocalc/scanner"
	"testing"
)

func TestIsBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "1 + 2",
			expected: true,
		},
		{
			input:    "1",
			expected: false,
		},
		{
			input:    "-1",
			expected: false,
		},
		{
			input:    "-1 + 2",
			expected: true,
		},
		{
			input:    "1 + 2 * 3",
			expected: true,
		},
		{
			input:    "(1 + 2)",
			expected: false,
		},
		{
			input:    "(1 + 2) * 3",
			expected: true,
		},
		{
			input:    "sin(1 + (2 - 3))",
			expected: false,
		},
		{
			input:    "PI",
			expected: false,
		},
		{
			input:    "PI * 4",
			expected: true,
		},
		{
			input:    "I = 12",
			expected: true,
		},
		{
			input:    "1.2",
			expected: false,
		},
		{
			input:    "1..2",
			expected: true,
		},
	}

	for i, test := range tests {
		expr, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(%d) scanner error => %s", i+1, err)
		}

		is := isBinary(expr)
		if is != test.expected {
			t.Errorf("(%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}
