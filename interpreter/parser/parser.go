package parser

import (
	"fmt"
	"gocalc/interpreter/lexemes"
	"gocalc/interpreter/scanner"
	"slices"
)

func parseRange(expr []scanner.Token) (NodeRange, error) {
	node := NodeRange{}

	if len(expr) < 3 {
		return node, fmt.Errorf("not enough tokens to parse expression as range")
	}

	// start

	start, ok := expr[0].(scanner.TokenNumber)
	if !ok || !start.IsInt() {
		return node, fmt.Errorf("expected a range start (an integer)")
	}

	node.Start = int(start.Value)

	// range operator

	oper, ok := expr[1].(scanner.TokenOperator)
	if !ok || oper.Value != lexemes.OPERATOR_RANGE {
		return node, fmt.Errorf("expected a range operator")
	}

	// end

	end, ok := expr[2].(scanner.TokenNumber)
	if !ok || !end.IsInt() {
		return node, fmt.Errorf("expected a range end (an integer)")
	}

	node.End = int(end.Value)

	return node, nil
}

// Checks if given expression (list of tokens) is binary (at least one
// top-level operator).
func isBinary(expr []scanner.Token) bool {
	depth := 0

	for _, t := range expr {
		_, isOper := t.(scanner.TokenOperator)
		if isOper && depth == 0 {
			return true
		} else {
			tPunc, ok := t.(scanner.TokenPunctuation)
			if ok {
				if tPunc.Value == lexemes.PUNCTUATION_LPAREN {
					depth++
				} else if tPunc.Value == lexemes.PUNCTUATION_RPAREN && depth > 0 {
					depth--
				}
			}
		}
	}

	return false
}

// Parses given expression as binary. If expression is not binary, returns
// [ErrNotABinaryExpression] error.
func parseBinary(expr []scanner.Token) (NodeBinary, error) {
	oper := scanner.TokenOperator{}
	left := []scanner.Token{}
	right := []scanner.Token{}

	depth := 0
	isLeftRead := false

	for _, t := range expr {
		tOper, ok := t.(scanner.TokenOperator)
		if ok && depth == 0 {
			if isLeftRead {
				left = append(left, oper)
				left = append(left, right...)
				oper = tOper
				right = right[:0]
			} else {
				oper = tOper
				isLeftRead = true
				continue
			}
		} else {
			if isLeftRead {
				right = append(right, t)
			} else {
				left = append(left, t)
			}

			tPunc, ok := t.(scanner.TokenPunctuation)
			if ok {
				if tPunc.Value == lexemes.PUNCTUATION_LPAREN {
					depth++
				} else if tPunc.Value == lexemes.PUNCTUATION_RPAREN {
					if depth > 0 {
						depth--
					} else {
						return NodeBinary{}, ErrMismatchedParenthesis
					}
				}
			}
		}
	}

	if !isLeftRead {
		return NodeBinary{}, fmt.Errorf("not a binary expression")
	}

	leftParsed, err := parseExpression(left)
	if err != nil {
		return NodeBinary{}, err
	}

	rightParsed, err := parseExpression(right)
	if err != nil {
		return NodeBinary{}, err
	}

	return NodeBinary{
		Operator: oper.Value,
		Left:     leftParsed,
		Right:    rightParsed,
	}, nil
}

func isIRangeFunctionArg(arg []scanner.Token) bool {
	if len(arg) >= 3 {
		_, ok := arg[0].(scanner.TokenWord)
		if !ok {
			return false
		}

		argAss, ok := arg[1].(scanner.TokenOperator)
		if !ok || argAss.Value != lexemes.OPERATOR_ASS {
			return false
		}

		_, err := parseRange(arg[2:])
		if err != nil {
			return false
		}

		return true
	}

	return false
}

func parseIRangeFunctionArg(arg []scanner.Token) (NodeIRangeFuncMainArg, error) {
	node := NodeIRangeFuncMainArg{}

	if len(arg) < 3 {
		return node, fmt.Errorf("not enough tokens in IRANGE function argument")
	}

	// variable

	argVar, ok := arg[0].(scanner.TokenWord)
	if !ok {
		return node, fmt.Errorf("expected a variable")
	}
	if slices.Contains(lexemes.DEFINED_CONSTANTS, argVar.Value) {
		return node, ErrVarAsConst
	}

	node.Variable = NodeVariable{
		Name: argVar.Value,
	}

	// assign operator

	argAss, ok := arg[1].(scanner.TokenOperator)
	if !ok || argAss.Value != lexemes.OPERATOR_ASS {
		return node, fmt.Errorf("expected an assign operator")
	}

	// range

	argRangeNode, err := parseRange(arg[2:])
	if err != nil {
		return node, err
	}

	node.Range = argRangeNode

	return node, nil
}

func parseFunctionArgs(tl []scanner.Token) ([]Node, error) {
	args := []Node{}

	sliced, err := SliceTokenListByComma(tl)
	if err != nil {
		return args, err
	}

	for _, arg := range sliced {
		var parsedArg Node

		if isIRangeFunctionArg(arg) {
			parsedArg, err = parseIRangeFunctionArg(arg)
			if err != nil {
				return args, err
			}
		} else {
			parsedArg, err = parseExpression(ParenthesiseExpression(arg))
			if err != nil {
				return args, err
			}
		}

		args = append(args, parsedArg)
	}

	return args, nil
}

//

func parseExpression(expr []scanner.Token) (Node, error) {
	if len(expr) == 0 {
		return InvalidNode{}, fmt.Errorf("empty expression")
	}

	switch firstToken := expr[0].(type) {
	case scanner.TokenPunctuation:
		if firstToken.Value == lexemes.PUNCTUATION_LPAREN {
			if isBinary(expr) {
				return parseBinary(expr)
			} else {
				return parseExpression(ReadParentheses(expr))
			}
		}

	case scanner.TokenNumber:
		return NodeNumber{
			Value: firstToken.Value,
		}, nil

	case scanner.TokenWord:
		if len(expr) > 1 {
			switch followUpToken := expr[1].(type) {
			case scanner.TokenPunctuation:
				{
					// if it is a function
					if followUpToken.Value == lexemes.PUNCTUATION_LPAREN {
						args, err := parseFunctionArgs(ReadParentheses(expr))
						return NodeFuncCall{
							Name:      firstToken.Value,
							Arguments: args,
						}, err
					}
				}

			default:
				return InvalidNode{}, fmt.Errorf("word followed by unexpected token %s", followUpToken.Type())
			}

		} else {
			if slices.Contains(lexemes.DEFINED_CONSTANTS, firstToken.Value) {
				return NodeConstant{
					Name: firstToken.Value,
				}, nil
			} else {
				return NodeVariable{
					Name: firstToken.Value,
				}, nil
			}
		}
	}

	return InvalidNode{}, fmt.Errorf("failed to parse an expression")
}

func Parse(tl []scanner.Token) (NodeRoot, error) {
	normalised := UnUnaryExpression(tl)
	v, err := parseExpression(ParenthesiseExpression(normalised))

	return NodeRoot{
		Value: v,
	}, err
}
