package parser

import (
	"gocalc/interpreter/scanner"
	"slices"
	"testing"
)

func TestUnUnaryExpression(t *testing.T) {
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
			input:    "-1",
			expected: "0-1",
		},
		{
			input:    "1 + (2 - 3)",
			expected: "1+(2-3)",
		},
		{
			input:    "4 * (-5 / 6)",
			expected: "4*(0-5/6)",
		},
		{
			input:    "-(-1)",
			expected: "0-(0-1)",
		},
		{
			input:    "-PHI",
			expected: "0-PHI",
		},
		{
			input:    "-SQRT(9)",
			expected: "0-SQRT(9)",
		},
		{
			input:    "-ABS(-1)",
			expected: "0-ABS(0-1)",
		},
	}

	for i, test := range tests {
		tl, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(case no.%d) => Scan error : %s", i, err)
		}

		ununared := UnUnaryExpression(tl)
		ununaredStr := scanner.StringifyTokens(ununared)

		if ununaredStr != test.expected {
			t.Errorf("(case no.%d) => got %s, expected %s", i, ununaredStr, test.expected)
		}
	}
}

func TestParenthesiseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "((((1))))",
		},
		{
			input:    "1 + 2",
			expected: "((((1)))+(((2))))",
		},
		{
			input:    "1 - 2",
			expected: "((((1)))-(((2))))",
		},
		{
			input:    "1 * 2",
			expected: "((((1))*((2))))",
		},
		{
			input:    "1 / 2",
			expected: "((((1))/((2))))",
		},
		{
			input:    "1^2",
			expected: "((((1)^(2))))",
		},
		{
			input:    "1 + (2 - 3)",
			expected: "((((1)))+(((((((2)))-(((3))))))))",
		},
		{
			input:    "1 + PI",
			expected: "((((1)))+(((PI))))",
		},
		{
			input:    "1.2 * COS(3)",
			expected: "((((1.2))*((COS((((3))))))))",
		},
		{
			input:    "12e3 / SUM(I=1..3, I)",
			expected: "((((12000))/((SUM((((I=1..3,I))))))))",
		},
	}

	for i, test := range tests {
		tl, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(case no.%d) => Scan error : %s", i, err)
		}

		parenthesised := ParenthesiseExpression(tl)
		parenthesisedStringified := scanner.StringifyTokens(parenthesised)

		if parenthesisedStringified != test.expected {
			t.Errorf("(case no.%d) => got %s, expected %s", i, parenthesisedStringified, test.expected)
		}
	}
}

func TestSliceTokenListByComma(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "1",
			expected: []string{"1"},
		},
		{
			input:    "1 + 2",
			expected: []string{"1+2"},
		},
		{
			input:    "1 + 2, 3 + 4",
			expected: []string{"1+2", "3+4"},
		},
		{
			input:    "1 + (2, 3) + 4",
			expected: []string{"1+(2,3)+4"},
		},
		{
			input:    "1 + (2 - 3), 4 * 5, 6",
			expected: []string{"1+(2-3)", "4*5", "6"},
		},
	}

	for i, test := range tests {
		tl, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(case no.%d) => Scan error : %s", i, err)
		}

		sliced, err := SliceTokenListByComma(tl)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}

		slicedAsStr := []string{}
		for _, sl := range sliced {
			slicedAsStr = append(slicedAsStr, scanner.StringifyTokens(sl))
		}

		isEqual := slices.Equal(slicedAsStr, test.expected)
		if !isEqual {
			t.Errorf("(case no.%d) => not equal : %q != %q", i, slicedAsStr, test.expected)
		}
	}
}
