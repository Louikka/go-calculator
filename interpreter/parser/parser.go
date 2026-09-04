package parser

import (
	"fmt"
	"gocalc/interpreter/lexemes"
	"gocalc/interpreter/scanner"
	"slices"
)

// Reads first encountered parentheses in expression. If no parentheses
// present, returns empty slice.
func readParentheses(expr []scanner.Token) []scanner.Token {
	depth := 0
	outExpr := []scanner.Token{}

	for _, t := range expr {
		tPunc, ok := t.(scanner.TokenPunctuation)
		if ok {
			if tPunc.Value == lexemes.PUNCTUATION_LPAREN {
				depth++
				if depth == 1 {
					continue
				}
			} else if tPunc.Value == lexemes.PUNCTUATION_RPAREN && depth > 0 {
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

type FunctionType int

const (
	// Function F(X) where F - name of the function, X - expression.
	FUNCTYPE_DEFAULT FunctionType = iota
	// Function F(I=N..M, X) where F - name of the function, I - name of
	// the variable, N - range start (inc.), M - range end (inc.), X -
	// expression (where the use of variable I are permitted).
	FUNCTYPE_IRANGE
)

func parseIRangeFunctionMainArg(arg []scanner.Token) (NodeIRangeFuncMainArg, error) {
	node := NodeIRangeFuncMainArg{}

	if len(arg) < 5 {
		return node, fmt.Errorf("not enough tokens in main argument of IRANGE function")
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

func parseFunctionArgs(tl []scanner.Token, funcType FunctionType) ([]Node, error) {
	args := []Node{}

	sliced, err := SliceTokenListByComma(tl)
	if err != nil {
		return args, err
	}

	switch funcType {
	case FUNCTYPE_DEFAULT:
		for _, expr := range sliced {
			parsedExpr, err := parseExpression(ParenthesiseExpression(expr))
			if err != nil {
				return args, err
			}

			args = append(args, parsedExpr)
		}

	case FUNCTYPE_IRANGE:
		{
			argsLen := len(sliced)
			if argsLen != 2 {
				return args, fmt.Errorf("expected 2 arguments, but got %d", argsLen)
			}

			mainArgExpr := sliced[0]
			secondaryArgExpr := sliced[1]

			mainArg, err := parseIRangeFunctionMainArg(mainArgExpr)
			if err != nil {
				return args, err
			}

			args = append(args, mainArg)

			secondaryArg, err := parseExpression(ParenthesiseExpression(secondaryArgExpr))
			if err != nil {
				return args, err
			}

			args = append(args, secondaryArg)
		}

	default:
		return args, fmt.Errorf("undefined funcType %d", funcType)
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
				return parseExpression(readParentheses(expr))
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
						funcName := firstToken.Value
						funcType := FUNCTYPE_DEFAULT

						if lexemes.IsIRangeFunction(funcName) {
							funcType = FUNCTYPE_IRANGE
						}

						args, err := parseFunctionArgs(readParentheses(expr), funcType)

						return NodeFuncCall{
							Name:      funcName,
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
