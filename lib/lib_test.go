package lib

import "testing"

func TestLongestStringLenInSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{},
			expected: 0,
		},
		{
			input:    []string{""},
			expected: 0,
		},
		{
			input:    []string{"", "a"},
			expected: 1,
		},
		{
			input:    []string{"", "a", "abc", "ab"},
			expected: 3,
		},
	}

	for i, test := range tests {
		l := LongestStringLenInSlice(test.input)
		if l != test.expected {
			t.Errorf("(case no.%d) => expected %d, got %d", i, test.expected, l)
		}
	}
}
