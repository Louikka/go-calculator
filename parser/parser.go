package parser

import (
	"fmt"
	"gocalc/lexer"
)

/* Types and constants *******************************************************/

type NodeType string

const (
	NODE_TYPE_ROOT       NodeType = "ROOT"
	NODE_TYPE_NUMBER     NodeType = "NUMBER"
	NODE_TYPE_CONSTANT   NodeType = "CONSTANT"
	NODE_TYPE_FUNCTION   NodeType = "FUNCTION"
	NODE_TYPE_EXPRESSION NodeType = "EXPRESSION"
	NODE_TYPE_BINARY     NodeType = "BINARY"
)

type Node struct {
	Type  NodeType
	Value any
}

type NodeValueNumber struct {
	Value float64
}
type NodeValueConstant struct {
	Name string
}
type NodeValueFunction struct {
	Name     string
	Argument Node
}
type NodeValueBinary struct {
	Operator byte
	Left     Node
	Right    Node
}

/* Helpers *******************************************************************/

func repeatInSlice[T any](item T, count uint) []T {
	// Pre-allocate the slice with a specific length
	s := make([]T, count)

	for i := range s {
		s[i] = item
	}

	return s
}

func parenthesizeExpression(input []lexer.Token) []lexer.Token {
	parenOpen := lexer.Token{
		Type:  lexer.TOKEN_TYPE_PUNCTUATION,
		Value: byte('('),
	}
	parenClose := lexer.Token{
		Type:  lexer.TOKEN_TYPE_PUNCTUATION,
		Value: byte(')'),
	}
	caret := lexer.Token{
		Type:  lexer.TOKEN_TYPE_OPERATOR,
		Value: byte('^'),
	}
	star := lexer.Token{
		Type:  lexer.TOKEN_TYPE_OPERATOR,
		Value: byte('*'),
	}
	slash := lexer.Token{
		Type:  lexer.TOKEN_TYPE_OPERATOR,
		Value: byte('/'),
	}
	plus := lexer.Token{
		Type:  lexer.TOKEN_TYPE_OPERATOR,
		Value: byte('+'),
	}
	minus := lexer.Token{
		Type:  lexer.TOKEN_TYPE_OPERATOR,
		Value: byte('-'),
	}

	output := []lexer.Token{}

	output = append(output, repeatInSlice(parenOpen, 4)...)

	for i, v := range input {
		t := v

		if t.Type == lexer.TOKEN_TYPE_OPERATOR || t.Type == lexer.TOKEN_TYPE_PUNCTUATION {
			switch t.Value.(byte) {
			case '(':
				output = append(output, repeatInSlice(parenOpen, 4)...)
				continue
			case ')':
				output = append(output, repeatInSlice(parenClose, 4)...)
				continue
			case '^':
				output = append(output, parenClose, caret, parenOpen)
				continue
			case '*':
				output = append(output, parenClose, parenClose, star, parenOpen, parenOpen)
				continue
			case '/':
				output = append(output, parenClose, parenClose, slash, parenOpen, parenOpen)
				continue
			case '+':
				//fmt.Println(i, v)
				// unary check: either first or had an operator expecting secondary argument
				if i == 0 || input[i-1].Type == lexer.TOKEN_TYPE_PUNCTUATION {
					output = append(output, plus)
				} else {
					output = append(output, parenClose, parenClose, parenClose, plus, parenOpen, parenOpen, parenOpen)
				}

				continue
			case '-':
				if i == 0 || input[i-1].Type == lexer.TOKEN_TYPE_PUNCTUATION {
					output = append(output, minus)
				} else {
					output = append(output, parenClose, parenClose, parenClose, minus, parenOpen, parenOpen, parenOpen)
				}

				continue
			}
		}

		output = append(output, t)
	}

	output = append(output, repeatInSlice(parenClose, 4)...)

	return output
}

func readParentheses(expr []lexer.Token) []lexer.Token {
	depth := 0
	inParenExpr := []lexer.Token{}

	for _, t := range expr {
		if t.Value == byte('(') {
			depth++
			if depth == 1 {
				continue
			}
		} else if t.Value == byte(')') && depth > 0 {
			depth--
			if depth == 0 {
				break
			}
		}

		if depth > 0 {
			inParenExpr = append(inParenExpr, t)
			continue
		}
	}

	return inParenExpr
}

func parseExpressionNode(expr []lexer.Token) (Node, error) {
	if len(expr) == 0 {
		return Node{}, fmt.Errorf("Empty expression (probably missing an operand).")
	}

	if expr[0].Value == byte('(') {
		// check if there is another binary expression on the same level
		pre := parseBinary(expr)
		if pre.Type == NODE_TYPE_BINARY {
			preValue := pre.Value.(_NodeValueBinaryPreparsed)

			__left, err := parseExpressionNode(preValue.Left)
			if err != nil {
				return Node{}, err
			}
			__right, err := parseExpressionNode(preValue.Right)
			if err != nil {
				return Node{}, err
			}

			return Node{
				Type: NODE_TYPE_BINARY,
				Value: NodeValueBinary{
					Operator: preValue.Operator,
					Left:     __left,
					Right:    __right,
				},
			}, nil
		}

		read := readParentheses(expr)
		bin := parseBinary(read)

		if bin.Type == NODE_TYPE_BINARY {
			binValue := bin.Value.(_NodeValueBinaryPreparsed)

			__left, err := parseExpressionNode(binValue.Left)
			if err != nil {
				return Node{}, err
			}
			__right, err := parseExpressionNode(binValue.Right)
			if err != nil {
				return Node{}, err
			}

			return Node{
				Type: NODE_TYPE_BINARY,
				Value: NodeValueBinary{
					Operator: binValue.Operator,
					Left:     __left,
					Right:    __right,
				},
			}, nil
		} else {
			return parseExpressionNode(bin.Value.([]lexer.Token))
		}

	} else if expr[0].Type == lexer.TOKEN_TYPE_NUMBER {
		return Node{
			Type:  NODE_TYPE_NUMBER,
			Value: expr[0].Value,
		}, nil

	} else if expr[0].Type == lexer.TOKEN_TYPE_CONSTANT {
		return Node{
			Type: NODE_TYPE_CONSTANT,
			Value: NodeValueConstant{
				Name: expr[0].Value.(string),
			},
		}, nil

	} else if expr[0].Type == lexer.TOKEN_TYPE_FUNCTION {
		__arg, err := parseExpressionNode(readParentheses(expr))
		if err != nil {
			return Node{}, err
		}

		return Node{
			Type: NODE_TYPE_FUNCTION,
			Value: NodeValueFunction{
				Name:     expr[0].Value.(string),
				Argument: __arg,
			},
		}, nil
	}

	return Node{}, nil
}

type _NodeValueBinaryPreparsed struct {
	Operator byte
	Left     []lexer.Token
	Right    []lexer.Token
}

func parseBinary(expr []lexer.Token) Node {
	var oper byte = 0
	left := []lexer.Token{}
	right := []lexer.Token{}

	depth := 0
	isLeftRead := false

	for _, t := range expr {
		if isLeftRead {
			right = append(right, t)
			continue
		}

		if t.Type == lexer.TOKEN_TYPE_OPERATOR && depth == 0 {
			oper = t.Value.(byte)
			isLeftRead = true
			continue
		} else {
			left = append(left, t)
			if t.Value == byte('(') {
				depth++
			} else if t.Value == byte(')') && depth > 0 {
				depth--
			}
			continue
		}
	}

	if isLeftRead {
		return Node{
			Type: NODE_TYPE_BINARY,
			Value: _NodeValueBinaryPreparsed{
				Operator: oper,
				Left:     left,
				Right:    right,
			},
		}
	} else {
		return Node{
			Type:  NODE_TYPE_EXPRESSION,
			Value: left,
		}
	}
}

/* Parser main ***************************************************************/

func Parse(tl []lexer.Token) (Node, error) {
	v, err := parseExpressionNode(parenthesizeExpression(tl))
	if err != nil {
		return Node{}, err
	}

	return Node{
		Type:  NODE_TYPE_ROOT,
		Value: v,
	}, nil
}
