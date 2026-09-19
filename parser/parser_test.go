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
		{
			input: "PI - 3.14",
			expected: NodeBinary{
				Operator: "-",
				Left: NodeIdentifier{
					Name: "PI",
				},
				Right: NodeNumber{
					Value: 3.14,
				},
			},
		},
		{
			input: "I = 12 / 4",
			expected: NodeBinary{
				Operator: "=",
				Left: NodeIdentifier{
					Name: "I",
				},
				Right: NodeBinary{
					Operator: "/",
					Left: NodeNumber{
						Value: 12,
					},
					Right: NodeNumber{
						Value: 4,
					},
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

func TestParseFuncArgs(t *testing.T) {
	tests := []struct {
		input    string
		expected []Node
	}{
		{
			input:    "",
			expected: []Node{},
		},
		{
			input: "1",
			expected: []Node{
				NewNodeNumber(1),
			},
		},
		{
			input: "1, 2",
			expected: []Node{
				NewNodeNumber(1),
				NewNodeNumber(2),
			},
		},
		{
			input: "1 + 2, 3 * 4",
			expected: []Node{
				NewNodeBinary(
					"+",
					NewNodeNumber(1),
					NewNodeNumber(2),
				),
				NewNodeBinary(
					"*",
					NewNodeNumber(3),
					NewNodeNumber(4),
				),
			},
		},
		{
			input: "1 + SOME_FUNC(2, 3) * 4",
			expected: []Node{
				NewNodeBinary(
					"+",
					NewNodeNumber(1),
					NewNodeBinary(
						"*",
						NewNodeFuncCall("SOME_FUNC", []Node{
							NewNodeNumber(2),
							NewNodeNumber(3),
						}),
						NewNodeNumber(4),
					),
				),
			},
		},
		{
			input: "1 + 2, SIN(3)",
			expected: []Node{
				NewNodeBinary(
					"+",
					NewNodeNumber(1),
					NewNodeNumber(2),
				),
				NewNodeFuncCall("SIN", []Node{
					NewNodeNumber(3),
				}),
			},
		},
	}

	for i, test := range tests {
		expr, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(%d) scanner error => %s", i+1, err)
		}

		args, err := parseFuncArgs(expr)
		if err != nil {
			t.Errorf("(%d) error => %s", i+1, err)
		}
		if !reflect.DeepEqual(args, test.expected) {
			t.Errorf("(%d) arguments are not matching => %+v", i+1, args)
		}
	}
}

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected Node
	}{
		{
			input:    "1",
			expected: NewNodeNumber(1),
		},
		{
			input: "1 + 2",
			expected: NewNodeBinary(
				"+",
				NewNodeNumber(1),
				NewNodeNumber(2),
			),
		},
		{
			input:    "PI",
			expected: NewNodeIdentifier("PI"),
		},
		{
			input:    "F ( )",
			expected: NewNodeFuncCall("F", []Node{}),
		},
	}

	for i, test := range tests {
		expr, err := scanner.Scan(test.input)
		if err != nil {
			t.Errorf("(%d) scanner error => %s", i+1, err)
		}

		node, err := parseExpression(expr)
		if err != nil {
			t.Errorf("(%d) error => %s", i+1, err)
		}
		if !reflect.DeepEqual(node, test.expected) {
			t.Errorf("(%d) nodes not matching => %+v", i+1, node)
		}
	}
}
