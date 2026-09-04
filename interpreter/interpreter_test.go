package interpreter

import (
	"math"
	"strconv"
	"testing"
)

func toString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func TestEvaluateString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1 + 2",
			expected: "3",
		},
		{
			input:    "1 - 2 + 3",
			expected: "2",
		},
		{
			input:    "1 - 2 + 3 - 4",
			expected: "-2",
		},
		{
			input:    "1 * 2 / 3 * 4",
			expected: toString(1.0 * 2.0 / 3.0 * 4.0),
		},
		{
			input:    "1.2 + 3.4",
			expected: "4.6",
		},
		{
			input:    "(10 - 5.4)",
			expected: "4.6",
		},
		{
			input:    "2 * 46",
			expected: "92",
		},
		{
			input:    "51 / 3",
			expected: "17",
		},
		{
			input:    "2 ^ 7",
			expected: "128",
		},
		{
			input:    "2 - 3 * 4",
			expected: "-10",
		},
		{
			input:    "2 * 3 - 4",
			expected: "2",
		},
		{
			input:    "PI * 3",
			expected: toString(math.Pi * 3),
		},
		{
			input:    "ABS(-12.5)",
			expected: "12.5",
		},
		{
			input:    "-ABS(-7e-2)",
			expected: "-0.07",
		},
		{
			input:    "abs(cos(pi))",
			expected: "1",
		},
		{
			input:    "ABS(SIN(3 * pi / 2))",
			expected: "1",
		},
		{
			input:    "SIN(pi) + COS (PI) - 1",
			expected: "-2",
		},
		{
			input:    "(SQRT(9) - ABS(-3 + 1)) * (9e3 + 999)",
			expected: "9999",
		},
		{
			input:    "-1 + 2",
			expected: "1",
		},
		{
			input:    "-1 / 4",
			expected: "-0.25",
		},
		{
			input:    "-1 / 4 + 0.25",
			expected: "0",
		},
		{
			input:    "-(1 + 2)",
			expected: "-3",
		},
		{
			input:    "-(-1)",
			expected: "1",
		},
	}

	for i, test := range tests {
		res, err := EvaluateString(test.input)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}

		toStr := toString(res)
		if toStr != test.expected {
			t.Errorf("(case no.%d) => expected %s, got %s", i, test.expected, toStr)
		}
	}
}
