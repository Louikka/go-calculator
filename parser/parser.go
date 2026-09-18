package parser

import (
	"fmt"
	"gocalc/scanner"
)

func parseBinary(expr []scanner.Token) (NodeBinary, error) {
	var operator scanner.TokenOperator
	left := []scanner.Token{}
	right := []scanner.Token{}

	r := false
	depth := 0

	for _, prec := range scanner.PossibleOperatorsPrecedence {
		for _, t := range expr {
			tOper, isOper := t.(scanner.TokenOperator)
			if isOper && depth == 0 && tOper.Precedence == prec {
				operator = tOper
				r = true
				continue
			}

			tPunc, isPunc := t.(scanner.TokenPunctuation)
			if isPunc {
				switch tPunc.Value {
				case "(":
					depth++

				case ")":
					depth--
					if depth < 0 {
						return NodeBinary{}, fmt.Errorf("mismatched parenthesis")
					}
				}
			}

			if r {
				right = append(right, t)
			} else {
				left = append(left, t)
			}
		}

		if r {
			break
		} else {
			r = false
			left = []scanner.Token{}
			right = []scanner.Token{}
		}
	}
}
