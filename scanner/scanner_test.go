package scanner

import (
	"testing"
)

func TestScannerInitialization(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedS string
	}{
		{
			name:      "just a number",
			input:     "0",
			expectedS: "0",
		},
		{
			name:      "just an expression",
			input:     "1 + 2",
			expectedS: "1 + 2",
		},
		{
			name:      "trimming leading and trailing whitespaces",
			input:     "  	\r 3 - 4     	\n	",
			expectedS: "3 - 4",
		},
		{
			name:      "capitalisation of the input",
			input:     "pi",
			expectedS: "PI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)

			if scanner.s != tt.expectedS {
				t.Errorf("strings mismatched (got %q, expected %q)", scanner.s, tt.expectedS)
			}
			if scanner.pos != 0 {
				t.Errorf("wrong initial position (got %d instead of 0).", scanner.pos)
			}
		})
	}
}

func TestScanner_IsLast(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			s:        "123",
			expected: false,
		},
		{
			s:        "12",
			expected: false,
		},
		{
			s:        "1",
			expected: true,
		},
		{
			s:        "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)

			is := scanner.isLast()
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestScanner_IsEnd(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			s:        "123",
			expected: false,
		},
		{
			s:        "12",
			expected: false,
		},
		{
			s:        "1",
			expected: false,
		},
		{
			s:        "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)

			is := scanner.isEnd()
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestScanner_Peek(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		pos          int
		peekArg      int
		expectedChar byte
	}{
		{
			s:            "abc",
			expectedChar: 'A',
		},
		{
			s:            "abc",
			pos:          1,
			expectedChar: 'B',
		},
		{
			s:            "abc",
			pos:          1,
			peekArg:      1,
			expectedChar: 'C',
		},
		{
			s:            "abc123",
			pos:          5,
			peekArg:      -2,
			expectedChar: '1',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)
			scanner.pos = tt.pos

			char := scanner.peek(tt.peekArg)
			if char != tt.expectedChar {
				t.Errorf("mismatched char %d instead of %d", char, tt.expectedChar)
			}
		})
	}
}

func TestScanner_Next(t *testing.T) {
	// todo
}

func TestScanner_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "just numbers",
			s:        "123",
			expected: false,
		},
		{
			name:     "numbers with some leading/trailing whitespaces",
			s:        "  123    ",
			expected: false,
		},
		{
			name:     "empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "one whitespace",
			s:        " ",
			expected: true,
		},
		{
			name:     "a bunch of whitespace characters",
			s:        " 	\r\n",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)

			is := scanner.IsEmpty()
			if is != tt.expected {
				t.Errorf("got %t, expected %t", is, tt.expected)
			}
		})
	}
}

func TestScanner_ReadNumber(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			s:        "1",
			expected: "1",
		},
		{
			s:        "1.2",
			expected: "1.2",
		},
		{
			s:        "3e4",
			expected: "30000",
		},
		{
			s:        "5e-6",
			expected: "0.000005",
		},
		{
			s:        "1e",
			expected: "1",
		},
		{
			s:        "1.",
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)

			n, err := scanner.readNumber()
			if err != nil {
				t.Errorf("error => %s", err)
			}

			n_str := n.String()
			if n_str != tt.expected {
				t.Errorf("got %s, expected %s", n_str, tt.expected)
			}
		})
	}
}

func TestScanner_ReadWord(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		expected     string
		expectedKind string
	}{
		{
			s:            "PI",
			expected:     "PI",
			expectedKind: WORD_KIND_IDENTIFIER,
		},
		{
			s:            "e",
			expected:     "E",
			expectedKind: WORD_KIND_IDENTIFIER,
		},
		{
			s:            "SQRT()",
			expected:     "SQRT",
			expectedKind: WORD_KIND_FUNCTION,
		},
		{
			s:            "ATAN ( )",
			expected:     "ATAN",
			expectedKind: WORD_KIND_FUNCTION,
		},
		{
			s:            "a1",
			expected:     "A1",
			expectedKind: WORD_KIND_IDENTIFIER,
		},
		{
			s:            "ABC123 * 4.5",
			expected:     "ABC123",
			expectedKind: WORD_KIND_IDENTIFIER,
		},
		{
			s:            "AB_C",
			expected:     "AB_C",
			expectedKind: WORD_KIND_IDENTIFIER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.s)

			w, err := scanner.readWord()
			if err != nil {
				t.Errorf("error => %s", err)
			}

			w_str := w.String()
			if w_str != tt.expected {
				t.Errorf("got %q, expected %q", w_str, tt.expected)
			}

			if w.Kind != tt.expectedKind {
				t.Errorf("(kind) got %q, expected %q", w.Kind, tt.expectedKind)
			}
		})
	}
}

func TestScannerOutputErrors(t *testing.T) {
	tests := []string{
		"0",
		"1 + 2",
		"1 + 2 * 3",
		"(1 + 2) * 3",
		"4 - PI",
		"SIN(5) / 6",
		"7 ^ 8",
		"ABS(-9)",
	}

	for _, test := range tests {
		_, err := Scan(test)
		if err != nil {
			t.Errorf("(%q) error => %s", test, err)
		}
	}
}
