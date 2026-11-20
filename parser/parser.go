package parser

import (
	"fmt"
	"gocalc/scanner"
)

func parseUnary(expr []scanner.Token) (NodeUnary, error) {
	if len(expr) == 0 {
		return NodeUnary{}, ErrEmptyExpr
	}

	uOper, isOper := expr[0].(scanner.TokenOperator)
	if isOper {
		exprParsed, err := parseExpression(expr[1:])
		if err != nil {
			return NodeUnary{}, err
		}

		return NewNodeUnary(uOper.Value, exprParsed), nil
	}

	lastI := len(expr) - 1
	uOper, isOper = expr[lastI].(scanner.TokenOperator)
	if isOper {
		exprParsed, err := parseExpression(expr[:lastI])
		if err != nil {
			return NodeUnary{}, err
		}

		return NewNodeUnary(uOper.Value, exprParsed), nil
	}

	return NodeUnary{}, ErrNotAnUnary
}

func parseBinary(expr []scanner.Token) (NodeBinary, error) {
	if len(expr) == 0 {
		return NodeBinary{}, ErrEmptyExpr
	}

	var err error = nil

	for _, prec := range scanner.SortOperatorsByPrecedence() {
		var operator scanner.TokenOperator
		left := []scanner.Token{}
		right := []scanner.Token{}

		isLeftRead := false
		depth := 0

		for i, t := range expr {
			tOper, isOper := t.(scanner.TokenOperator)
			if isOper && i == 0 { // skip unary operator
				left = append(left, t)
				continue
			}
			if isOper && depth == 0 && tOper.Precedence == prec {
				if isLeftRead {
					left = append(left, operator)
					left = append(left, right...)
					operator = tOper
					right = right[:0]
				} else {
					operator = tOper
					isLeftRead = true
				}

				continue // skip operator token
			} else {
				depth, err = checkDepth(t, depth)
				if err != nil {
					return NodeBinary{}, err
				}
			}

			if isLeftRead {
				right = append(right, t)
			} else {
				left = append(left, t)
			}
		}

		if isLeftRead {
			leftParsed, err := parseExpression(left)
			if err != nil {
				return NodeBinary{}, err
			}

			rightParsed, err := parseExpression(right)
			if err != nil {
				return NodeBinary{}, err
			}

			return NewNodeBinary(operator.Value, leftParsed, rightParsed), nil
		}
	}

	return NodeBinary{}, ErrNotABinary
}

func parseFuncArgs(expr []scanner.Token) ([]Node, error) {
	es, err := sliceTokenListByComma(expr)
	args := []Node{}

	if err != nil {
		return args, nil
	}

	for _, e := range es {
		parsed, err := parseExpression(e)
		if err != nil {
			return args, err
		}

		args = append(args, parsed)
	}

	return args, nil
}

func parseExpression(expr []scanner.Token) (Node, error) {
	if len(expr) == 0 {
		return NodeInvalid{}, ErrEmptyExpr
	}

	if isBinary(expr) {
		return parseBinary(expr)
	}

	if isUnary(expr) {
		return parseUnary(expr)
	}

	switch t := expr[0].(type) {
	case scanner.TokenNumber:
		return NewNodeNumber(t.Value), nil

	case scanner.TokenWord:
		if t.Kind == scanner.WORD_KIND_FUNCTION {
			args, err := parseFuncArgs(readParentheses(expr))
			if err != nil {
				return NodeInvalid{}, err
			}

			return NewNodeFuncCall(t.Value, args), nil
		} else {
			return NewNodeIdentifier(t.Value), nil
		}

	case scanner.TokenOperator:
		templ := "unexpected operator %q at the start of the " +
			"expression (did you forget to un-unary the expression?)"
		panic(fmt.Sprintf(templ, t.Value))

	case scanner.TokenPunctuation:
		if t.Value == "(" {
			return parseExpression(readParentheses(expr))
		} else {
			return NodeInvalid{}, fmt.Errorf("unexpected token %q", t.Value)
		}
	}

	return NodeInvalid{}, fmt.Errorf("failed to parse an expression")
}

//---------------------------------------------------------------------------//

func Parse(input []scanner.Token) (NodeRoot, error) {
	node, err := parseExpression(input)
	if err != nil {
		return NodeRoot{}, err
	}

	return NodeRoot{
		Expression: node,
	}, nil
}
