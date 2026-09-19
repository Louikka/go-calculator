package parser

import (
	"fmt"
	"gocalc/scanner"
)

func parseBinary(expr []scanner.Token) (NodeBinary, error) {
	for _, prec := range scanner.OperatorsPrecedenceSorted {
		var operator scanner.TokenOperator
		left := []scanner.Token{}
		right := []scanner.Token{}

		isLeftRead := false
		depth := 0

		for _, t := range expr {
			if isLeftRead {
				right = append(right, t)
			} else {
				var err error = nil

				tOper, isOper := t.(scanner.TokenOperator)
				if isOper && depth == 0 && tOper.Precedence == prec {
					operator = tOper
					isLeftRead = true
					continue // skip operator token
				}

				depth, err = checkDepth(t, depth)
				if err != nil {
					return NodeBinary{}, err
				}

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
		return NodeInvalid{}, fmt.Errorf("empty expression")
	}

	if isBinary(expr) {
		return parseBinary(expr)
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

	case scanner.TokenPunctuation:
		if t.Value == "(" {
			return parseExpression(readParentheses(expr))
		} else {
			return NodeInvalid{}, fmt.Errorf("unexpected token \"%s\"", t.Value)
		}
	}

	return NodeInvalid{}, fmt.Errorf("failed to parse an expression")
}

//---------------------------------------------------------------------------//

func Parse(input []scanner.Token) (NodeRoot, error) {
	node, err := parseExpression(input)
	return NodeRoot{
		Expression: node,
	}, err
}
