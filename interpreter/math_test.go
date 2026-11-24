package interpreter

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{
			name:     "calling function with no arguments should return 0",
			values:   []float64{},
			expected: 0,
		},
		{
			name:     "calling function with 1 argument should return same argument",
			values:   []float64{1},
			expected: 1,
		},
		{
			values:   []float64{1, 2},
			expected: 3,
		},
		{
			values:   []float64{0.1, 0.2},
			expected: 0.3,
		},
		{
			values:   []float64{0, 1, 2, 3, 5, 10},
			expected: 21,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Add(test.values...)
			if n != test.expected {
				t.Errorf("%g != %g", n, test.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{
			name:     "calling function with no arguments should return 0",
			values:   []float64{},
			expected: 0,
		},
		{
			name:     "calling function with 1 argument should return same argument",
			values:   []float64{1},
			expected: 1,
		},
		{
			values:   []float64{1, 2},
			expected: -1,
		},
		{
			values:   []float64{0.1, 0.2},
			expected: -0.1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Sub(test.values...)
			if n != test.expected {
				t.Errorf("%g != %g", n, test.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{
			name:     "calling function with no arguments should return 0",
			values:   []float64{},
			expected: 0,
		},
		{
			name:     "calling function with 1 argument should return same argument",
			values:   []float64{1},
			expected: 1,
		},
		{
			values:   []float64{1, 2},
			expected: 2,
		},
		{
			values:   []float64{0.1, 0.2},
			expected: 0.02,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Mul(test.values...)
			if n != test.expected {
				t.Errorf("%g != %g", n, test.expected)
			}
		})
	}
}
