package interpreter

import (
	"gocalc/parser"
	"math"
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
