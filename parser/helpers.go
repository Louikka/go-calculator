package parser

import (
	"fmt"
	"gocalc/scanner"
)

// Checks given token, and if it parenthesis, returns new depth. If token is
// not a parenthesis, returns unchanged depth. Also checks for parenthesis
// matching and returns [ErrMismatchedParenthesis] error if depth less than 0.
func checkDepth(t scanner.Token, depth int) (int, error) {
	tPunc, ok := t.(scanner.TokenPunctuation)
	if ok {
		switch tPunc.Value {
		case "(":
			return depth + 1, nil

		case ")":
			depth--
			if depth < 0 {
				return depth, ErrMismatchedParenthesis
			} else {
				return depth, nil
			}
		}
	}

	return depth, nil
}

// Reads first encountered parentheses in expression. If no parentheses
// present, returns empty slice.
func readParentheses(expr []scanner.Token) []scanner.Token {
	outExpr := []scanner.Token{}

	depth := 0

	for _, t := range expr {
		tPunc, ok := t.(scanner.TokenPunctuation)
		if ok {
			if tPunc.Value == "(" {
				depth++
				if depth == 1 {
					continue
				}
			} else if tPunc.Value == ")" && depth > 0 {
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

// Divides list of tokens by top-level commas.
func sliceTokenListByComma(tl []scanner.Token) ([][]scanner.Token, error) {
	groups := [][]scanner.Token{}

	depth := 0

	if len(tl) == 0 {
		return groups, nil
	} else {
		groups = append(groups, []scanner.Token{})
	}

	for _, t := range tl {
		tPunc, ok := t.(scanner.TokenPunctuation)
		if ok {
			switch tPunc.Value {
			case ",":
				if depth == 0 {
					groups = append(groups, []scanner.Token{})
					continue
				}

			case "(":
				depth++

			case ")":
				if depth > 0 {
					depth--
				} else {
					return groups, ErrMismatchedParenthesis
				}
			}
		}

		groupsLastIndex := len(groups) - 1
		groups[groupsLastIndex] = append(groups[groupsLastIndex], t)
	}

	return groups, nil
}

// Converts unary operations to binary (e.g. "-1" to "0 - 1").
func unUnaryExpression(expr []scanner.Token) ([]scanner.Token, error) {
	out := []scanner.Token{}

	tryAppend := func(oper scanner.TokenOperator) error {
		if oper.CanBeUnary() {
			out = append(out, scanner.NewTokenNumber(0))
			return nil
		} else {
			return fmt.Errorf("operator %q cannot be unary", oper.Value)
		}
	}

	for i, t := range expr {
		oper, isOper := t.(scanner.TokenOperator)
		if isOper {
			if i == 0 {
				// first token in list
				err := tryAppend(oper)
				if err != nil {
					return out, err
				}

			} else /* i > 0 */ {
				punc, isPunc := expr[i-1].(scanner.TokenPunctuation)
				if isPunc && punc.Value == "(" {
					// first token after left parenthesis
					err := tryAppend(oper)
					if err != nil {
						return out, err
					}
				}
			}
		}

		out = append(out, t)
	}

	return out, nil
}
