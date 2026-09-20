package parser

import (
	"gocalc/scanner"
	"slices"
	"testing"
)

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

		sliced, err := sliceTokenListByComma(tl)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}

		slicedAsStr := []string{}
		for _, sl := range sliced {
			slicedAsStr = append(slicedAsStr, scanner.Stringify(sl, ""))
		}

		isEqual := slices.Equal(slicedAsStr, test.expected)
		if !isEqual {
			t.Errorf("(case no.%d) => not equal : %q != %q", i, slicedAsStr, test.expected)
		}
	}
}

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
			input:    "+1",
			expected: "0+1",
		},
		{
			input:    "-1 + 2",
			expected: "0-1+2",
		},
		{
			input:    "-(1 + 2)",
			expected: "0-(1+2)",
		},
		{
			input:    "(-1) + 2",
			expected: "(0-1)+2",
		},
		{
			input:    "-ABS(-4)",
			expected: "0-ABS(0-4)",
		},
	}

	for i, test := range tests {
		tl, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(%d) => Scan error : %s", i+1, err)
		}

		expr, err := unUnaryExpression(tl)
		if err != nil {
			t.Errorf("(%d) => error : %s", i+1, err)
		}

		str := scanner.Stringify(expr, "")
		if str != test.expected {
			t.Errorf("(%d) => %s != %s", i+1, str, test.expected)
		}
	}
}
