package parser

import (
	"fmt"
	"gocalc/interpreter/scanner"
)

// Converts unary operations to binary
func unUnaryExpression(expr []scanner.Token) []scanner.Token {
	out := []scanner.Token{}

	for i, t := range expr {
		tIsOper := t.Type == scanner.TOKEN_TYPE_OPERATOR
		if tIsOper && (t.Value == scanner.OPERATOR_ADD || t.Value == scanner.OPERATOR_SUB) {
			if i == 0 {
				out = append(out, scanner.NewNumberToken(0))
			} else /* i > 0 */ {
				prevT := expr[i-1]
				prevTIsPunc := prevT.Type == scanner.TOKEN_TYPE_PUNCTUATION
				if prevTIsPunc && prevT.Value == scanner.PUNCTUATION_LPAREN {
					out = append(out, scanner.NewNumberToken(0))
				}
			}
		}

		out = append(out, t)
	}

	return out
}

// Reads expression in parentheses. If no parentheses present in expression, returns empty slice.
func readParentheses(expr []scanner.Token) []scanner.Token {
	depth := 0
	outExpr := []scanner.Token{}

	for _, t := range expr {
		if t.Type == scanner.TOKEN_TYPE_PUNCTUATION {
			if t.Value == scanner.PUNCTUATION_LPAREN {
				depth++
				if depth == 1 {
					continue
				}
			} else if t.Value == scanner.PUNCTUATION_RPAREN && depth > 0 {
				depth--
				if depth == 0 {
					break
				}
			}
		}

		if depth > 0 {
			outExpr = append(outExpr, t)
			continue
		}
	}

	return outExpr
}

func parenthesizeExpression(input []scanner.Token) []scanner.Token {
	// https://en.wikipedia.org/wiki/Operator-precedence_parser#Full_parenthesization

	tAdd := scanner.NewOperatorToken(scanner.OPERATOR_ADD)
	tSub := scanner.NewOperatorToken(scanner.OPERATOR_SUB)
	tMul := scanner.NewOperatorToken(scanner.OPERATOR_MUL)
	tDiv := scanner.NewOperatorToken(scanner.OPERATOR_DIV)
	tPow := scanner.NewOperatorToken(scanner.OPERATOR_POW)
	tLParen := scanner.NewPunctuationToken(scanner.PUNCTUATION_LPAREN)
	tRParen := scanner.NewPunctuationToken(scanner.PUNCTUATION_RPAREN)

	output := []scanner.Token{}

	output = append(output, tLParen, tLParen, tLParen, tLParen)

	for i, t := range input {
		if t.Type == scanner.TOKEN_TYPE_OPERATOR || t.Type == scanner.TOKEN_TYPE_PUNCTUATION {
			switch t.Value {
			case scanner.OPERATOR_POW:
				output = append(output, tRParen, tPow, tLParen)

			case scanner.OPERATOR_MUL:
				output = append(output, tRParen, tRParen, tMul, tLParen, tLParen)

			case scanner.OPERATOR_DIV:
				output = append(output, tRParen, tRParen, tDiv, tLParen, tLParen)

			case scanner.OPERATOR_ADD:
				// unary check: either first or had an operator expecting secondary argument
				if i == 0 || input[i-1].Value == scanner.PUNCTUATION_LPAREN {
					output = append(output, tAdd)
				} else {
					output = append(output, tRParen, tRParen, tRParen, tAdd, tLParen, tLParen, tLParen)
				}

			case scanner.OPERATOR_SUB:
				// unary check
				if i == 0 || input[i-1].Value == scanner.PUNCTUATION_LPAREN {
					output = append(output, tSub)
				} else {
					output = append(output, tRParen, tRParen, tRParen, tSub, tLParen, tLParen, tLParen)
				}

			case scanner.PUNCTUATION_LPAREN:
				output = append(output, tLParen, tLParen, tLParen, tLParen)

			case scanner.PUNCTUATION_RPAREN:
				output = append(output, tRParen, tRParen, tRParen, tRParen)
			}
		} else {
			output = append(output, t)
		}
	}

	output = append(output, tRParen, tRParen, tRParen, tRParen)

	return output
}

type _NodeValueBinary struct {
	Operator string
	Left     []scanner.Token
	Right    []scanner.Token
}

func parseExpressionNode(expr []scanner.Token) (Node, error) {
	if len(expr) == 0 {
		return Node{}, fmt.Errorf("Empty expression (probably missing an operand).")
	}

	if expr[0].Value == scanner.PUNCTUATION_LPAREN {
		// check if there is another binary expression on the same level
		pre := parseBinary(expr)
		if pre.Type == NODE_TYPE_BINARY {
			preValue := pre.Value.(_NodeValueBinary)

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
			binValue := bin.Value.(_NodeValueBinary)

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
			return parseExpressionNode(bin.Value.([]scanner.Token))
		}

	} else if expr[0].Type == scanner.TOKEN_TYPE_NUMBER {
		return Node{
			Type:  NODE_TYPE_NUMBER,
			Value: expr[0].Value,
		}, nil

	} else if expr[0].Type == scanner.TOKEN_TYPE_VARIABLE {
		return Node{
			Type: NODE_TYPE_VARIABLE,
			Value: NodeValueVariable{
				Name: expr[0].Value.(string),
			},
		}, nil

	} else if expr[0].Type == scanner.TOKEN_TYPE_CONSTANT {
		return Node{
			Type: NODE_TYPE_CONSTANT,
			Value: NodeValueConstant{
				Name: expr[0].Value.(string),
			},
		}, nil

	} else if expr[0].Type == scanner.TOKEN_TYPE_FUNCTION {
		__arg, err := parseExpressionNode(readParentheses(expr))
		if err != nil {
			return Node{}, err
		}

		return Node{
			Type: NODE_TYPE_FUNCTION_CALL,
			Value: NodeValueFunction{
				Name:     expr[0].Value.(string),
				Argument: __arg,
			},
		}, nil
	}

	return Node{}, nil
}

func parseBinary(expr []scanner.Token) Node {
	var oper string = ""
	left := []scanner.Token{}
	right := []scanner.Token{}

	depth := 0
	isLeftRead := false

	for _, t := range expr {
		if isLeftRead {
			right = append(right, t)
			continue
		}

		if t.Type == scanner.TOKEN_TYPE_OPERATOR && depth == 0 {
			oper = t.Value.(string)
			isLeftRead = true
			continue
		} else {
			left = append(left, t)
			if t.Value == scanner.PUNCTUATION_LPAREN {
				depth++
			} else if t.Value == scanner.PUNCTUATION_RPAREN && depth > 0 {
				depth--
			}
			continue
		}
	}

	if isLeftRead {
		return Node{
			Type: NODE_TYPE_BINARY,
			Value: _NodeValueBinary{
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

func Parse(tl []scanner.Token) (Node, error) {
	v, err := parseExpressionNode(parenthesizeExpression(unUnaryExpression(tl)))
	if err != nil {
		return Node{}, err
	}

	return Node{
		Type:  NODE_TYPE_ROOT,
		Value: v,
	}, nil
}
