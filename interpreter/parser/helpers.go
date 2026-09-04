package parser

import (
	"gocalc/interpreter/lexemes"
	"gocalc/interpreter/scanner"
)

// Checks if operator token can be unary.
func canBeUnaryOper(t scanner.TokenOperator) bool {
	return t.Value == lexemes.OPERATOR_ADD || t.Value == lexemes.OPERATOR_SUB
}

// Converts unary operations to binary (e.g. "-1" to "0 - 1").
func UnUnaryExpression(expr []scanner.Token) []scanner.Token {
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
					if ok && prevTPunc.Value == lexemes.PUNCTUATION_LPAREN {
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

// https://en.wikipedia.org/wiki/Operator-precedence_parser#Full_parenthesization
func ParenthesiseExpression(input []scanner.Token) []scanner.Token {

	tLParen := scanner.NewTokenPunctuation(lexemes.PUNCTUATION_LPAREN)
	tRParen := scanner.NewTokenPunctuation(lexemes.PUNCTUATION_RPAREN)

	output := []scanner.Token{}

	output = append(output, tLParen, tLParen, tLParen, tLParen)

	for _, t := range input {
		switch tok := t.(type) {
		case scanner.TokenOperator:
			{
				switch tok.Value {
				case lexemes.OPERATOR_POW:
					output = append(output, tRParen, tok, tLParen)

				case lexemes.OPERATOR_MUL:
					output = append(output, tRParen, tRParen, tok, tLParen, tLParen)

				case lexemes.OPERATOR_DIV:
					output = append(output, tRParen, tRParen, tok, tLParen, tLParen)

				case lexemes.OPERATOR_ADD:
					output = append(output, tRParen, tRParen, tRParen, tok, tLParen, tLParen, tLParen)

				case lexemes.OPERATOR_SUB:
					output = append(output, tRParen, tRParen, tRParen, tok, tLParen, tLParen, tLParen)

				default:
					output = append(output, tok)
				}
			}

		case scanner.TokenPunctuation:
			{
				switch tok.Value {
				case lexemes.PUNCTUATION_LPAREN:
					output = append(output, tLParen, tLParen, tLParen, tLParen)

				case lexemes.PUNCTUATION_RPAREN:
					output = append(output, tRParen, tRParen, tRParen, tRParen)

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

// Divides list of tokens by top-level commas.
func SliceTokenListByComma(tl []scanner.Token) ([][]scanner.Token, error) {
	groups := [][]scanner.Token{}

	depth := 0

	groups = append(groups, []scanner.Token{})

	for _, t := range tl {
		tPunc, ok := t.(scanner.TokenPunctuation)
		if ok {
			if tPunc.Value == lexemes.PUNCTUATION_COMMA && depth == 0 {
				groups = append(groups, []scanner.Token{})
				continue
			}

			if tPunc.Value == lexemes.PUNCTUATION_LPAREN {
				depth++
			} else if tPunc.Value == lexemes.PUNCTUATION_RPAREN {
				if depth > 0 {
					depth--
				} else {
					return groups, ErrMismatchedParenthesis
				}
			}
		}

		groups[len(groups)-1] = append(groups[len(groups)-1], t)
	}

	return groups, nil
}
