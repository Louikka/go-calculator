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
