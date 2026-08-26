package interpreter

import (
	"math"
	"testing"
)

type _TestSpec struct {
	input          string
	expectedResult float64
}

func TestInterpreterEvaluationFromString(t *testing.T) {
	tests := []_TestSpec{
		{
			input:          "1 + 2",
			expectedResult: 3,
		},
		{
			input:          "1.2 + 3.4",
			expectedResult: 4.6,
		},
		{
			input:          "10 - 5.4",
			expectedResult: 4.6,
		},
		{
			input:          "2 * 46",
			expectedResult: 92,
		},
		{
			input:          "51 / 3",
			expectedResult: 17,
		},
		{
			input:          "2 ^ 7",
			expectedResult: 128,
		},
		{
			input:          "PI * 3",
			expectedResult: 9.425,
		},
		{
			input:          "ABS(-12.5)",
			expectedResult: 12.5,
		},
	}

	for _, test := range tests {
		res, err := EvaluateString(test.input)
		if err != nil {
			t.Errorf("error in string evaluation => \"%s\".", err)
		}

		rounded := math.Round(res*1000) / 1000
		if test.expectedResult != rounded {
			t.Errorf("string evaluation outputs are not equal => %f != %f.", test.expectedResult, rounded)
		}
	}
}
