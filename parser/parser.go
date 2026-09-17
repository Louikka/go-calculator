package parser

import (
	"gocalc/lib"
	token "gocalc/tokens"
)

func isTopALeftParenthesis(stack lib.Stack[token.Token]) bool {
	top := stack.Top()

	punc, isPunc := top.(token.TokenPunctuation)
	if isPunc && punc.IsLeftParenthesis() {
		return true
	}

	return false
}

// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func ToPostfix(expr []token.Token) ([]token.Token, error) {
	output := lib.NewStack[token.Token]()
	stack := lib.NewStack[token.Token]()

	for _, t := range expr {
		switch tt := t.(type) {
		case token.TokenNumber:
			output.Append(tt)

		case token.TokenWord:
			if tt.Kind == token.WORD_KIND_FUNCTION {
				stack.Append(tt)
			} else {
				output.Append(tt)
			}

		case token.TokenOperator:
			for !stack.IsEmpty() {
				o1 := tt
				o2, isOper := stack.Top().(token.TokenOperator)
				if !isOper {
					break
				}

				c1 := o2.Precedence > o1.Precedence
				c2 := o1.Precedence == o2.Precedence
				c3 := o1.Associativity == "LEFT"

				if c1 || (c2 && c3) {
					output.Append(o2)
					stack.Pop()
				} else {
					break
				}
			}

			stack.Append(tt)

		case token.TokenPunctuation:
			switch tt.Value {
			case ",":
				for !stack.IsEmpty() {
					if isTopALeftParenthesis(stack) {
						break
					}

					output.Append(stack.Pop())
				}

			case "(":
				stack.Append(tt)

			case ")":
				for !stack.IsEmpty() {
					if isTopALeftParenthesis(stack) {
						break
					}

					output.Append(stack.Pop())
				}

				if stack.IsEmpty() {
					return output.Stack, ErrMismatchedParenthesis
				}

				if !isTopALeftParenthesis(stack) {
					return output.Stack, ErrMismatchedParenthesis
				}

				// discard left parenthesis from stack top
				stack.Pop()

				if !stack.IsEmpty() {
					tTop := stack.Top()
					tTopWord, isWord := tTop.(token.TokenWord)
					if isWord && (tTopWord.Kind == token.WORD_KIND_FUNCTION) {
						output.Append(tTop)
						stack.Pop()
					}
				}

			default:
				stack.Append(tt)
			}

		default:
			output.Append(tt)
		}
	}

	for !stack.IsEmpty() {
		if isTopALeftParenthesis(stack) {
			return output.Stack, ErrMismatchedParenthesis
		}

		output.Append(stack.Pop())
	}

	return output.Stack, nil
}
