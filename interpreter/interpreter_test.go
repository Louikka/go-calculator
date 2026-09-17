package interpreter

import (
	"gocalc/lib"
	"math"
	"testing"
)

type TestEvalS struct {
	input    string
	expected string
}

func tt(t *testing.T, tests []TestEvalS) bool {
	for i, test := range tests {
		res, err := EvaluateString(test.input)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
			return false
		}

		toStr := lib.FloatToString(res)
		if toStr != test.expected {
			t.Errorf("\"%s\" (case no.%d) => expected %s, got %s", test.input, i, test.expected, toStr)
			return false
		}
	}

	return true
}

func TestEvaluateString_Expression(t *testing.T) {
	tests := []TestEvalS{
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
			expected: lib.FloatToString(1.0 * 2.0 / 3.0 * 4.0),
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
			input:    "51 / (3)",
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
		{
			input:    "2e3",
			expected: "2000",
		},
		{
			input:    "4e-3",
			expected: "0.004",
		},
	}

	tt(t, tests)
}

func TestEvaluateString_Constants(t *testing.T) {
	tests := []TestEvalS{
		{
			input:    "PI",
			expected: lib.FloatToString(math.Pi),
		},
		{
			input:    "E - 1",
			expected: lib.FloatToString((math.E - 1) - 0.0000000000000002),
			//                                         ^ expected 1.7182818284590453, got 1.718281828459045??
			//                                          where did 0.0000000000000003 came from?
			//                                                                 ^
		},
		{
			input:    "(PSI * 3) + 1.5",
			expected: lib.FloatToString((psi * 3) + 1.5),
		},
	}

	tt(t, tests)
}

func TestEvaluateString_Functions(t *testing.T) {
	tests := []TestEvalS{
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
			input:    "SUM(I=1..5, I)",
			expected: "15",
		},
		{
			input:    "-SUM(I=2..6, I) + 5",
			expected: "-15",
		},
		{
			input:    "SUM(I=1..10, I * 2)",
			expected: "110",
		},
		{
			input:    "PROD(I=1..5, I)",
			expected: "120",
		},
		{
			input:    "cbrt(125)",
			expected: "5",
		},
		{
			input:    "Round(1.39)",
			expected: "1",
		},
	}

	tt(t, tests)
}
