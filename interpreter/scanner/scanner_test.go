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

	for _, test := range tests {
		scanner := _NewScanner(test.input)

		if scanner.s != test.expected {
			t.Errorf("\"%s\" => strings mismatched (got \"%s\", expected \"%s\")", test.input, scanner.s, test.expected)
		}
		if scanner.pos != 0 {
			t.Errorf("\"%s\" => wrong initial position (got %d instead of 0).", test.input, scanner.pos)
		}
	}
}

func TestScannerCoreFunctions(t *testing.T) {
	const test_str = "123"
	const test_strlen = len(test_str)

	scanner := _NewScanner(test_str)

	peekedChar, err := scanner.peek(0)
	if err != nil {
		t.Errorf("peek() => %s", err)
	}
	if peekedChar != test_str[0] {
		t.Errorf("peek() => (peeking at pos 0) mismatched char (\"%c\" instead of \"%c\")", test_str[0], peekedChar)
	}

	nextChar, err := scanner.next()
	if err != nil {
		t.Errorf("next() => %s", err)
	}
	if nextChar != test_str[1] {
		t.Errorf("next() => mismatched char (\"%c\" instead of \"%c\")", test_str[1], nextChar)
	}

	peekedChar, err = scanner.peek(1)
	if err != nil {
		t.Errorf("peek() => %s", err)
	}
	if peekedChar != test_str[2] {
		t.Errorf("peek() => (peeking at pos 2) mismatched char (\"%c\" instead of \"%c\")", test_str[2], peekedChar)
	}

	// set position to the last
	scanner.pos = test_strlen - 1

	if !scanner.isLast() {
		t.Errorf("isLast() => current position %d in string with length of %d", scanner.pos, test_strlen)
	}

	scanner.pos++

	if !scanner.isEnd() {
		t.Errorf("isEnd() => current position %d in string with length of %d", scanner.pos, test_strlen)
	}

	scanner = _NewScanner("abcdefg123456789")
	for i := 0; !scanner.isEnd(); i++ {
		_, err := scanner.peek(0)
		if err != nil {
			t.Errorf("iteration failed => peek() error at iteration %d : %s", i, err)
		}

		if !scanner.isLast() {
			_, err = scanner.next()
			if err != nil {
				t.Errorf("iteration failed => next() error at iteration %d : %s", i, err)
			}
		} else {
			break
		}
	}
}

func TestScannerReadKeyword(t *testing.T) {
	tests := []struct {
		s      string
		ttype  TokenType
		tvalue string
	}{
		{
			s:      "PI",
			ttype:  TOKEN_TYPE_CONSTANT,
			tvalue: "PI",
		},
		{
			s:      "E",
			ttype:  TOKEN_TYPE_CONSTANT,
			tvalue: "E",
		},
		{
			s:      "NONEXISTINGCONSTANT",
			ttype:  TOKEN_TYPE_CONSTANT,
			tvalue: "NONEXISTINGCONSTANT",
		},
		{
			s:      "COS(12.3)",
			ttype:  TOKEN_TYPE_FUNCTION,
			tvalue: "COS",
		},
		{
			s:      "ABS()",
			ttype:  TOKEN_TYPE_FUNCTION,
			tvalue: "ABS",
		},
		{
			s:      "NONEXISTINGFUNCTION()",
			ttype:  TOKEN_TYPE_FUNCTION,
			tvalue: "NONEXISTINGFUNCTION",
		},
	}

	for _, test := range tests {
		scanner := _NewScanner(test.s)

		token, err := scanner.readKeyword()
		if err != nil {
			t.Errorf("\"%s\" => %s", test.s, err)
		}

		if token.Type != test.ttype {
			t.Errorf("\"%s\" => token type mismatched (got \"%s\", expected \"%s\")", test.s, token.Type, test.ttype)
		}
		if token.Value != test.tvalue {
			t.Errorf("\"%s\" => token value mismatched (got \"%s\", expected \"%s\")", test.s, token.Value, test.tvalue)
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
