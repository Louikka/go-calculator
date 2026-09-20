package parser

import "gocalc/scanner"

// Checks if given expression (list of tokens) is binary (at least one
// top-level operator).
func isBinary(expr []scanner.Token) bool {
	depth := 0

	for i, t := range expr {
		_, isOper := t.(scanner.TokenOperator)
		if isOper {
			if i == 0 {
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
