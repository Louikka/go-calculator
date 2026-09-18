package parser

import (
	"gocalc/scanner"
	"reflect"
	"testing"
)

func TestParseBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected NodeBinary
	}{
		{
			input: "1 + 2",
			expected: NodeBinary{
				Operator: "+",
				Left: NodeNumber{
					Value: 1,
				},
				Right: NodeNumber{
					Value: 2,
				},
			},
		},
		{
			input: "1 + 2 * 3",
			expected: NodeBinary{
				Operator: "+",
				Left: NodeNumber{
					Value: 1,
				},
				Right: NodeBinary{
					Operator: "*",
					Left: NodeNumber{
						Value: 2,
					},
					Right: NodeNumber{
						Value: 3,
					},
				},
			},
		},
		{
			input: "(1 + 2) * 3",
			expected: NodeBinary{
				Operator: "*",
				Left: NodeBinary{
					Operator: "+",
					Left: NodeNumber{
						Value: 1,
					},
					Right: NodeNumber{
						Value: 2,
					},
				},
				Right: NodeNumber{
					Value: 3,
				},
			},
		},
	}

	for i, test := range tests {
		expr, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(%d) scanner error => %s", i+1, err)
		}

		node, err := parseBinary(expr)
		if err != nil {
			t.Errorf("(%d) error => %s", i+1, err)
		}
		if !reflect.DeepEqual(node, test.expected) {
			t.Errorf("(%d) nodes not matching => %+v", i+1, node)
		}
	}
}
