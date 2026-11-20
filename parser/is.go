package parser

import "gocalc/scanner"

// Checks if given expression (list of tokens) is binary (at least one
// top-level operator).
func isBinary(expr []scanner.Token) bool {
	depth := 0

	for i, t := range expr {
		_, isOper := t.(scanner.TokenOperator)
		if isOper {
			if i == 0 || i == len(expr)-1 {
				// unary operators are not counting
				continue
			}
			if depth == 0 {
				return true
			}
		} else {
			tPunc, isPunc := t.(scanner.TokenPunctuation)
			if isPunc {
				if tPunc.Value == "(" {
					depth++
				} else if tPunc.Value == ")" && depth > 0 {
					depth--
				}
			}
		}
	}

	return false
}

// Checks if given expression (list of tokens) is unary (not a binary and
// starts or ends with an operator).
func isUnary(expr []scanner.Token) bool {
	if isBinary(expr) {
		return false
	}

	if len(expr) > 0 {
		_, isOper := expr[0].(scanner.TokenOperator)
		if isOper {
			return true
		}

		_, isOper = expr[len(expr)-1].(scanner.TokenOperator)
		if isOper {
			return true
		}
	}

	return false
}
