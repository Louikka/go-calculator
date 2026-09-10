package parser

import (
	l "gocalc/interpreter/lexemes"
	"gocalc/interpreter/scanner"
)

// Checks if operator token can be unary.
func canBeUnaryOper(t scanner.TokenOperator) bool {
	return t.Value == l.OPERATOR_ADD || t.Value == l.OPERATOR_SUB
}

// Converts unary operations to binary (e.g. "-1" to "0 - 1").
func unUnaryExpression(expr []scanner.Token) []scanner.Token {
	out := []scanner.Token{}

	for i, t := range expr {
		tOper, isOper := t.(scanner.TokenOperator)
		if isOper {
			if canBeUnaryOper(tOper) {
				if i == 0 {
					// first token in list
					out = append(out, scanner.NewTokenNumber(0))

				} else /* i > 0 */ {
					prevT := expr[i-1]
					prevTPunc, ok := prevT.(scanner.TokenPunctuation)
					if ok && prevTPunc.Value == l.PUNCTUATION_LPAREN {
						// first token after left parenthesis
						out = append(out, scanner.NewTokenNumber(0))
					}
				}
			}
		}

		out = append(out, t)
	}

	return out
}

func normalise(tl []scanner.Token) []scanner.Token {
	return unUnaryExpression(tl)
}

// Parentheses an expression (excluding function calls) -
// https://en.wikipedia.org/wiki/Operator-precedence_parser#Full_parenthesization
func parenthesiseExpression(input []scanner.Token) []scanner.Token {

	tLParen := scanner.NewTokenPunctuation(l.PUNCTUATION_LPAREN)
	tRParen := scanner.NewTokenPunctuation(l.PUNCTUATION_RPAREN)

	output := []scanner.Token{}

	output = append(output, tLParen, tLParen, tLParen, tLParen)

	isFunc := false
	funcDepth := 0

	for i, t := range input {
		switch tok := t.(type) {
		case scanner.TokenOperator:
			if !isFunc {
				switch tok.Value {
				case l.OPERATOR_POW:
					output = append(output, tRParen, tok, tLParen)

				case l.OPERATOR_MUL:
					output = append(output, tRParen, tRParen, tok, tLParen, tLParen)

				case l.OPERATOR_DIV:
					output = append(output, tRParen, tRParen, tok, tLParen, tLParen)

				case l.OPERATOR_ADD:
					output = append(output, tRParen, tRParen, tRParen, tok, tLParen, tLParen, tLParen)

				case l.OPERATOR_SUB:
					output = append(output, tRParen, tRParen, tRParen, tok, tLParen, tLParen, tLParen)

				default:
					output = append(output, tok)
				}
			} else {
				output = append(output, tok)
			}

		case scanner.TokenPunctuation:
			{
				switch tok.Value {
				case l.PUNCTUATION_LPAREN:
					{
						if isFunc {
							funcDepth++
						}

						if i > 0 {
							// if previous token is a word (i.e function call)
							_, ok := input[i-1].(scanner.TokenWord)
							if ok {
								output = append(output, tok)
								isFunc = true
								continue
							}
						}

						output = append(output, tLParen, tLParen, tLParen, tLParen)
					}

				case l.PUNCTUATION_RPAREN:
					{
						if isFunc {
							if funcDepth > 0 {
								funcDepth--
							} else {
								output = append(output, tok)
								isFunc = false
							}
						} else {
							output = append(output, tRParen, tRParen, tRParen, tRParen)
						}
					}

				default:
					output = append(output, tok)
				}
			}

		default:
			output = append(output, tok)
		}
	}

	output = append(output, tRParen, tRParen, tRParen, tRParen)

	return output
}

// Reads first encountered parentheses in expression. If no parentheses
// present, returns empty slice.
func readParentheses(expr []scanner.Token) []scanner.Token {
	outExpr := []scanner.Token{}

	depth := 0

	for _, t := range expr {
		tPunc, ok := t.(scanner.TokenPunctuation)
		if ok {
			if tPunc.Value == l.PUNCTUATION_LPAREN {
				depth++
				if depth == 1 {
					continue
				}
			} else if tPunc.Value == l.PUNCTUATION_RPAREN && depth > 0 {
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
			case l.PUNCTUATION_COMMA:
				if depth == 0 {
					groups = append(groups, []scanner.Token{})
					continue
				}

			case l.PUNCTUATION_LPAREN:
				depth++

			case l.PUNCTUATION_RPAREN:
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
