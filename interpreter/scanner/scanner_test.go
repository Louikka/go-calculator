package scanner

import (
	"testing"
)

func TestScannerInitialization(t *testing.T) {
	const test_str = "123"

	scanner := _NewScanner(test_str)

	if scanner.s != test_str {
		t.Errorf("Strings not matched => \"%s\" != \"%s\".", test_str, scanner.s)
	}

	if scanner.pos != 0 {
		t.Errorf("Wrong current position initialization => %d.", scanner.pos)
	}
}

func TestScannerCoreFunctions(t *testing.T) {
	const test_str = "123"
	const test_strlen = len(test_str)

	scanner := _NewScanner(test_str)

	peekedChar, err := scanner.peek(0)
	if err != nil || peekedChar != test_str[0] {
		t.Errorf("Scanner.peek() error or peeking at wrong poition => error or mismatched char %c != %c.", test_str[0], peekedChar)
	}

	nextChar, err := scanner.next()
	if err != nil || nextChar != test_str[1] {
		t.Errorf("Scanner.next() error => error or mismatched char \"%c\" != \"%c\".", test_str[1], nextChar)
	}

	peekedChar, err = scanner.peek(1)
	if err != nil || peekedChar != test_str[2] {
		t.Errorf("Scanner.peek() error or peeking at wrong poition => error or mismatched char %c != %c.", test_str[2], peekedChar)
	}

	scanner.pos = test_strlen

	if !scanner.isEndOfString() {
		t.Errorf("Cannot determine if end of string => current position %d in string with length of %d.", scanner.pos, test_strlen)
	}

	scanner2 := _NewScanner("abcdefghjk")
	for i := 0; !scanner2.isEndOfString(); i++ {
		_, err := scanner2.peek(0)
		if err != nil {
			t.Errorf("Iteration failed => Scanner.peek() error at iteration %d.", i)
		}

		scanner2.next()
	}
}

func TestScannerOutput(t *testing.T) {
	_, err := Scan("1 + 2")
	if err != nil {
		t.Errorf("Error while Scan() => %v.", err)
	}
}
