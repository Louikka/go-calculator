package scanner

import "testing"

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "0",
			input:    '0',
			expected: true,
		},
		{
			name:     "1",
			input:    '1',
			expected: true,
		},
		{
			name:     "letter",
			input:    'a',
			expected: false,
		},
		{
			name:     "some unrelated symbol",
			input:    '@',
			expected: false,
		},
		{
			name:     "whitespace/tab",
			input:    '	',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isDigit(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "a",
			input:    'a',
			expected: true,
		},
		{
			name:     "B",
			input:    'B',
			expected: true,
		},
		{
			name:     "z",
			input:    'z',
			expected: true,
		},
		{
			name:     "whitespace",
			input:    ' ',
			expected: false,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "other symbols",
			input:    '&',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isLetter(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "space",
			input:    ' ',
			expected: true,
		},
		{
			name:     "tab",
			input:    '	',
			expected: true,
		},
		{
			name:     "vertical tab",
			input:    11,
			expected: true,
		},
		{
			name:     "new line character (\\n)",
			input:    '\n',
			expected: true,
		},
		{
			name:     "caret return character (\\r)",
			input:    '\r',
			expected: true,
		},
		{
			name:     "star symbol",
			input:    '*',
			expected: false,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "letter",
			input:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isWhitespace(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsWordStart(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "letter",
			input:    'a',
			expected: true,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "percent sign",
			input:    '%',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isWordStart(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsWordBody(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "letter",
			input:    'a',
			expected: true,
		},
		{
			name:     "number",
			input:    '1',
			expected: true,
		},
		{
			name:     "underscore",
			input:    '_',
			expected: true,
		},
		{
			name:     "space",
			input:    ' ',
			expected: false,
		},
		{
			name:     "hypen",
			input:    '-',
			expected: false,
		},
		{
			name:     "question mark",
			input:    '?',
			expected: false,
		},
		{
			name:     "dot",
			input:    '.',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isWordBody(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsOperatorStart(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "plus sign",
			input:    '+',
			expected: true,
		},
		{
			name:     "slash",
			input:    '/',
			expected: true,
		},
		{
			name:     "dot",
			input:    '.',
			expected: true,
		},
		{
			name:     "hash",
			input:    '#',
			expected: false,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "letter",
			input:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isOperatorStart(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestIsPunctuationStart(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "lparen",
			input:    '(',
			expected: true,
		},
		{
			name:     "comma",
			input:    ',',
			expected: true,
		},
		{
			name:     "dollar sign",
			input:    '$',
			expected: false,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "letter",
			input:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isPunctuationStart(tt.input)
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}
