package scanner

import "testing"

type lenHelper struct {
	s string
}

func (h lenHelper) Len() int {
	return len(h.s)
}

func TestLongestLen(t *testing.T) {
	tests := []struct {
		name     string
		input    []Lener
		expected int
	}{
		{
			input: []Lener{
				lenHelper{},
			},
			expected: 0,
		},
		{
			input: []Lener{
				lenHelper{s: ""},
				lenHelper{s: "a"},
				lenHelper{s: "ab"},
			},
			expected: 2,
		},
		{
			input: []Lener{
				lenHelper{s: "asdjk"},
				lenHelper{s: "12s1dfd"},
				lenHelper{s: "s"},
				lenHelper{s: "da1"},
			},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LongestLen(tt.input)
			if l != tt.expected {
				t.Errorf("got %d, expected %d", l, tt.expected)
			}
		})
	}
}

func TestStringify(t *testing.T) {
	tests := []struct {
		name     string
		input    []Token
		delim    string
		expected string
	}{
		{
			input: []Token{
				NewTokenNumber(0),
			},
			expected: "0",
		},
		{
			input: []Token{
				NewTokenNumber(1),
				NewTokenNumber(2),
				NewTokenNumber(3),
			},
			expected: "123",
		},
		{
			input: []Token{
				NewTokenNumber(1),
				NewTokenOperator("+"),
				NewTokenNumber(2),
			},
			delim:    " ",
			expected: "1 + 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stringify(tt.input, tt.delim)
			if s != tt.expected {
				t.Errorf("got %s, expected %s", s, tt.expected)
			}
		})
	}
}
