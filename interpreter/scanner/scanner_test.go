package scanner

import (
	"testing"
)

func TestScannerInitialization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "0",
			expected: "0",
		},
		{
			input:    " 1 + 2",
			expected: "1 + 2",
		},
		{
			input:    "  	 3 - 4     		",
			expected: "3 - 4",
		},
		{
			input:    "   pi * 10 ",
			expected: "PI * 10",
		},
		{
			input:    " Tan( 10 / 3 ) ",
			expected: "TAN( 10 / 3 )",
		},
	}

	for i, test := range tests {
		scanner := NewScanner(test.input)

		if scanner.s != test.expected {
			t.Errorf("(case no.%d) => strings mismatched (got \"%s\", expected \"%s\")", i, scanner.s, test.expected)
		}
		if scanner.pos != 0 {
			t.Errorf("(case no.%d) => wrong initial position (got %d instead of 0).", i, scanner.pos)
		}
	}
}

func TestScanner_IsLast(t *testing.T) {
	cases := []struct {
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

	for i, c := range cases {
		scanner := NewScanner(c.s)
		got := scanner.isLast()
		if got != c.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, c.expected, got)
		}
	}
}

func TestScanner_IsEnd(t *testing.T) {
	cases := []struct {
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

	for i, c := range cases {
		scanner := NewScanner(c.s)
		got := scanner.isEnd()
		if got != c.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, c.expected, got)
		}
	}
}

func TestScanner_Peek(t *testing.T) {
	const s = "123"

	cases := []struct {
		peekPos      int
		expectedChar byte
	}{
		{
			peekPos:      0,
			expectedChar: s[0],
		},
		{
			peekPos:      1,
			expectedChar: s[1],
		},
		{
			peekPos:      2,
			expectedChar: s[2],
		},
	}

	scanner := NewScanner(s)

	for i, c := range cases {
		char, err := scanner.peek(c.peekPos)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}
		if char != c.expectedChar {
			t.Errorf("(case no.%d) => mismatched char %d instead of %d", i, char, c.expectedChar)
		}
	}
}

func TestScanner_Next(t *testing.T) {
	const s = "123"

	expectedChars := []byte{
		s[1],
		s[2],
	}

	scanner := NewScanner(s)

	for i, expected := range expectedChars {
		char, err := scanner.next()
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}
		if char != expected {
			t.Errorf("(case no.%d) => mismatched char %d instead of %d", i, char, expected)
		}
	}
}

func TestScanner_IsEmpty(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{
			s:        "123",
			expected: false,
		},
		{
			s:        "  123    ",
			expected: false,
		},
		{
			s:        "",
			expected: true,
		},
		{
			s:        "   ",
			expected: true,
		},
		{
			s:        " 	\n",
			expected: true,
		},
	}

	for i, c := range cases {
		scanner := NewScanner(c.s)
		got := scanner.IsEmpty()
		if got != c.expected {
			t.Errorf("(case no.%d) => expected %t, got %t", i, c.expected, got)
		}
	}
}

func TestScanner_ReadNumber(t *testing.T) {
	tests := []struct {
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
			s:        "1.",
			expected: "1",
		},
		{
			s:        "1.a",
			expected: "1",
		},
	}

	for i, test := range tests {
		scanner := NewScanner(test.s)
		n, err := scanner.readNumber()
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
		}
		if n.ToString() != test.expected {
			t.Errorf("(case no.%d) => expected %s, got %s", i, test.expected, n.ToString())
		}
	}
}

func TestScannerOutput(t *testing.T) {
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
			t.Errorf("\"%s\" => %s", test, err)
		}
	}
}
