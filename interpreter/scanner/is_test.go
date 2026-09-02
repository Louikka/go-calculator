package scanner

import "testing"

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			input:    '1',
			expected: true,
		},
		{
			input:    '0',
			expected: true,
		},
		{
			input:    '9',
			expected: true,
		},
		{
			input:    'a',
			expected: false,
		},
		{
			input:    'T',
			expected: false,
		},
		{
			input:    '@',
			expected: false,
		},
		{
			input:    '	',
			expected: false,
		},
	}

	for i, test := range tests {
		is := isDigit(test.input)
		if is != test.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			input:    'a',
			expected: true,
		},
		{
			input:    'B',
			expected: true,
		},
		{
			input:    'z',
			expected: true,
		},
		{
			input:    ' ',
			expected: false,
		},
		{
			input:    '1',
			expected: false,
		},
		{
			input:    '&',
			expected: false,
		},
	}

	for i, test := range tests {
		is := isLetter(test.input)
		if is != test.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			input:    ' ',
			expected: true,
		},
		{
			input:    '	',
			expected: true,
		},
		{
			input:    11, // vertical tab
			expected: true,
		},
		{
			input:    '\n',
			expected: true,
		},
		{
			input:    '\r',
			expected: true,
		},
		{
			input:    '*',
			expected: false,
		},
		{
			input:    '1',
			expected: false,
		},
		{
			input:    'a',
			expected: false,
		},
	}

	for i, test := range tests {
		is := isWhitespace(test.input)
		if is != test.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}

func TestIsOperatorStart(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			input:    '+',
			expected: true,
		},
		{
			input:    '/',
			expected: true,
		},
		{
			input:    '.',
			expected: true,
		},
		{
			input:    '#',
			expected: false,
		},
		{
			input:    '1',
			expected: false,
		},
		{
			input:    'a',
			expected: false,
		},
	}

	for i, test := range tests {
		is := isOperatorStart(test.input)
		if is != test.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}

func TestIsPunctuationStart(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			input:    '(',
			expected: true,
		},
		{
			input:    ',',
			expected: true,
		},
		{
			input:    '$',
			expected: false,
		},
		{
			input:    '1',
			expected: false,
		},
		{
			input:    'a',
			expected: false,
		},
	}

	for i, test := range tests {
		is := isPunctuationStart(test.input)
		if is != test.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, test.expected, is)
		}
	}
}
