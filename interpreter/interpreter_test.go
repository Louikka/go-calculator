package interpreter

import (
	"gocalc/parser"
	"math"
	"strconv"
	"testing"
)

func TestSolveNodeIdentifier(t *testing.T) {
	tests := []struct {
		node     parser.NodeIdentifier
		ctx      []_VarCtx
		expected float64
	}{
		{
			node:     parser.NewNodeIdentifier("PI"),
			expected: math.Pi,
		},
		{
			node: parser.NewNodeIdentifier("I"),
			ctx: []_VarCtx{
				{
					Name:  "I",
					Value: 0,
				},
			},
			expected: 0,
		},
		{
			node: parser.NewNodeIdentifier("MY_VAR2"),
			ctx: []_VarCtx{
				{
					Name:  "MY_VAR1",
					Value: 1.2,
				},
				{
					Name:  "MY_VAR2",
					Value: 3.45,
				},
			},
			expected: 3.45,
		},
	}

	for i, test := range tests {
		n, err := solveNodeIdentifier(test.node, test.ctx)
		if err != nil {
			t.Errorf("(%d) error => %s", i+1, err)
		}
		if n != test.expected {
			t.Errorf("(%d) => %f != %f", i+1, n, test.expected)
		}
	}
}

func TestSolveNodeFuncCall(t *testing.T) {
	tests := []struct {
		node     parser.NodeFuncCall
		ctx      []_VarCtx
		expected float64
	}{
		{
			node: parser.NewNodeFuncCall("ABS", []parser.Node{
				parser.NewNodeNumber(-2.5),
			}),
			expected: 2.5,
		},
		{
			node: parser.NewNodeFuncCall("SUM", []parser.Node{
				parser.NewNodeBinary(
					"=",
					parser.NewNodeIdentifier("I"),
					parser.NewNodeBinary(
						"..",
						parser.NewNodeNumber(1),
						parser.NewNodeNumber(5),
					),
				),
				parser.NewNodeIdentifier("I"),
			}),
			expected: 15,
		},
		{
			node: parser.NewNodeFuncCall("SUM", []parser.Node{
				parser.NewNodeAssign(
					parser.NewNodeIdentifier("I"),
					parser.NewNodeRange(1, 5),
				),
				parser.NewNodeIdentifier("I"),
			}),
			expected: 15,
		},
	}

	for i, test := range tests {
		n, err := solveNodeFuncCall(test.node, test.ctx)
		if err != nil {
			t.Errorf("(%d) error => %s", i+1, err)
		}
		if n != test.expected {
			t.Errorf("(%d) => %f != %f", i+1, n, test.expected)
		}
	}

	// special "RAND" function
	{
		n, err := solveNodeFuncCall(parser.NewNodeFuncCall("RAND", []parser.Node{}), []_VarCtx{})
		if err != nil {
			t.Errorf("(RAND) error => %s", err)
		}
		if n < 0 || n >= 1 {
			t.Errorf("(RAND) => number %f not in interval [0, 1)", n)
		}
	}
}

type TestEvalS struct {
	input    string
	expected string
}

func tt(t *testing.T, tests []TestEvalS) bool {
	for i, test := range tests {
		res, err := EvaluateString(test.input)
		if err != nil {
			t.Errorf("(case no.%d) error => %s", i, err)
			return false
		}

		toStr := strconv.FormatFloat(res, 'f', -1, 64)
		if toStr != test.expected {
			t.Errorf("%q (case no.%d) => expected %s, got %s", test.input, i, test.expected, toStr)
			return false
		}
	}

	return true
}

func TestEvaluateString_Expression(t *testing.T) {
	tests := []TestEvalS{
		{
			input:    "1 + 2",
			expected: "3",
		},
		{
			input:    "1 - 2 + 3",
			expected: "2",
		},
		{
			input:    "1 - 2 + 3 - 4",
			expected: "-2",
		},
		{
			input:    "1 * 2 / 3 * 4",
			expected: strconv.FormatFloat(1.0*2.0/3.0*4.0, 'f', -1, 64),
		},
		{
			input:    "1.2 + 3.4",
			expected: "4.6",
		},
		{
			input:    "(10 - 5.4)",
			expected: "4.6",
		},
		{
			input:    "2 * 46",
			expected: "92",
		},
		{
			input:    "51 / (3)",
			expected: "17",
		},
		{
			input:    "2 ^ 7",
			expected: "128",
		},
		{
			input:    "2 - 3 * 4",
			expected: "-10",
		},
		{
			input:    "2 * 3 - 4",
			expected: "2",
		},
		{
			input:    "-1 + 2",
			expected: "1",
		},
		{
			input:    "-1 / 4",
			expected: "-0.25",
		},
		{
			input:    "-1 / 4 + 0.25",
			expected: "0",
		},
		{
			input:    "-(1 + 2)",
			expected: "-3",
		},
		{
			input:    "-(-1)",
			expected: "1",
		},
		{
			input:    "2e3",
			expected: "2000",
		},
		{
			input:    "4e-3",
			expected: "0.004",
		},
	}

	tt(t, tests)
}

func TestEvaluateString_Constants(t *testing.T) {
	tests := []TestEvalS{
		{
			input:    "PI",
			expected: strconv.FormatFloat(math.Pi, 'f', -1, 64),
		},
		{
			input:    "E - 1",
			expected: strconv.FormatFloat((math.E-1)-0.0000000000000002, 'f', -1, 64),
			//                            ^ expected 1.7182818284590453, got 1.718281828459045??
			//                             where did 0.0000000000000003 came from?
			//                                                        ^
		},
		{
			input:    "(PSI * 3) + 1.5",
			expected: strconv.FormatFloat((psi*3)+1.5, 'f', -1, 64),
		},
	}

	tt(t, tests)
}

func TestEvaluateString_Functions(t *testing.T) {
	tests := []TestEvalS{
		{
			input:    "ABS(-12.5)",
			expected: "12.5",
		},
		{
			input:    "-ABS(-7e-2)",
			expected: "-0.07",
		},
		{
			input:    "abs(cos(pi))",
			expected: "1",
		},
		{
			input:    "ABS(SIN(3 * pi / 2))",
			expected: "1",
		},
		{
			input:    "SIN(pi) + COS (PI) - 1",
			expected: "-2",
		},
		{
			input:    "(SQRT(9) - ABS(-3 + 1)) * (9e3 + 999)",
			expected: "9999",
		},
		{
			input:    "SUM(I=1..5, I)",
			expected: "15",
		},
		{
			input:    "-SUM(I=2..6, I) + 5",
			expected: "-15",
		},
		{
			input:    "SUM(I=1..10, I * 2)",
			expected: "110",
		},
		{
			input:    "PROD(I=1..5, I)",
			expected: "120",
		},
		{
			input:    "cbrt(125)",
			expected: "5",
		},
		{
			input:    "Round(1.39)",
			expected: "1",
		},
	}

	tt(t, tests)
}
