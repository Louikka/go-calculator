package parser

import (
	"gocalc/interpreter/scanner"
)

type Stack struct {
	Stack []scanner.Token
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Len() int {
	return len(s.Stack)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}

// Peeks at the top of the stack.
func (s *Stack) Top() scanner.Token {
	if len(s.Stack) > 0 {
		return s.Stack[len(s.Stack)-1]
	} else {
		return scanner.InvalidToken{}
	}
}

func (s *Stack) Append(t ...scanner.Token) {
	s.Stack = append(s.Stack, t...)
}

func (s *Stack) Pop() scanner.Token {
	if len(s.Stack) > 0 {
		i := len(s.Stack) - 1
		top := s.Stack[i]
		s.Stack = s.Stack[:i]
		return top
	} else {
		return scanner.InvalidToken{}
	}
}

func (s *Stack) IsTopALeftParenthesis() bool {
	top := s.Top()

	punc, isPunc := top.(scanner.TokenPunctuation)
	if isPunc && punc.IsLeftParenthesis() {
		return true
	}

	return false
}

// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func converExpressionToRPN(expr []scanner.Token) ([]scanner.Token, error) {
	output := NewStack()
	stack := NewStack()

	for _, t := range expr {
		switch tt := t.(type) {
		case scanner.TokenNumber:
			output.Append(tt)

		case scanner.TokenWord:
			if tt.Kind == scanner.WORD_KIND_FUNCTION {
				stack.Append(tt)
			} else {
				output.Append(tt)
			}

		case scanner.TokenOperator:
			for !stack.IsEmpty() {
				tTopOper, isOper := stack.Top().(scanner.TokenOperator)
				if !isOper {
					break
				}

				o1 := tt.Definition()
				o2 := tTopOper.Definition()

				if (o2.Precedence > o1.Precedence) || ((o1.Precedence == o2.Precedence) && o1.IsLeftAssociative) {
					output.Append(tTopOper)
					stack.Pop()
				} else {
					break
				}
			}

			stack.Append(tt)

		case scanner.TokenPunctuation:
			switch tt.Value {
			case ",":
				for !stack.IsEmpty() {
					if stack.IsTopALeftParenthesis() {
						break
					}

					output.Append(stack.Pop())
				}

			case "(":
				stack.Append(tt)

			case ")":
				for !stack.IsEmpty() {
					if stack.IsTopALeftParenthesis() {
						break
					}

					output.Append(stack.Pop())
				}

				if stack.IsEmpty() {
					return output.Stack, ErrMismatchedParenthesis
				}

				if !stack.IsTopALeftParenthesis() {
					return output.Stack, ErrMismatchedParenthesis
				}

				// discard left parenthesis from stack top
				stack.Pop()

				if !stack.IsEmpty() {
					tTop := stack.Top()
					tTopWord, isWord := tTop.(scanner.TokenWord)
					if isWord && (tTopWord.Kind == scanner.WORD_KIND_FUNCTION) {
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
		if stack.IsTopALeftParenthesis() {
			return output.Stack, ErrMismatchedParenthesis
		}

		output.Append(stack.Pop())
	}

	return output.Stack, nil
}

func parseRPN(expr []scanner.Token) (Node, error) {
	return InvalidNode{}, nil
}

/* ***************************************************************************/

// func parseRange(expr []scanner.Token) (NodeRange, error) {
// 	node := NodeRange{}

// 	if len(expr) < 3 {
// 		return node, fmt.Errorf("not enough tokens to parse expression as range")
// 	}

// 	// start

// 	start, ok := expr[0].(scanner.TokenNumber)
// 	if !ok || !start.IsInt() {
// 		return node, fmt.Errorf("expected a range start (an integer)")
// 	}

// 	node.Start = int(start.Value)

// 	// range operator

// 	oper, ok := expr[1].(scanner.TokenOperator)
// 	if !ok || oper.Value != l.OPERATOR_RANGE {
// 		return node, fmt.Errorf("expected a range operator")
// 	}

// 	// end

// 	end, ok := expr[2].(scanner.TokenNumber)
// 	if !ok || !end.IsInt() {
// 		return node, fmt.Errorf("expected a range end (an integer)")
// 	}

// 	node.End = int(end.Value)

// 	return node, nil
// }

// // Checks if given expression (list of tokens) is binary (at least one
// // top-level operator).
// func isBinary(expr []scanner.Token) bool {
// 	depth := 0

// 	for _, t := range expr {
// 		_, isOper := t.(scanner.TokenOperator)
// 		if isOper && depth == 0 {
// 			return true
// 		} else {
// 			tPunc, ok := t.(scanner.TokenPunctuation)
// 			if ok {
// 				if tPunc.Value == l.PUNCTUATION_LPAREN {
// 					depth++
// 				} else if tPunc.Value == l.PUNCTUATION_RPAREN && depth > 0 {
// 					depth--
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

// // Parses given expression as binary. If expression is not binary, returns
// // [ErrNotABinaryExpression] error.
// func parseBinary(expr []scanner.Token) (NodeBinary, error) {
// 	oper := scanner.TokenOperator{}
// 	left := []scanner.Token{}
// 	right := []scanner.Token{}

// 	depth := 0
// 	isLeftRead := false

// 	for _, t := range expr {
// 		tOper, ok := t.(scanner.TokenOperator)
// 		if ok && depth == 0 {
// 			if isLeftRead {
// 				left = append(left, oper)
// 				left = append(left, right...)
// 				oper = tOper
// 				right = right[:0]
// 			} else {
// 				oper = tOper
// 				isLeftRead = true
// 				continue
// 			}
// 		} else {
// 			if isLeftRead {
// 				right = append(right, t)
// 			} else {
// 				left = append(left, t)
// 			}

// 			tPunc, ok := t.(scanner.TokenPunctuation)
// 			if ok {
// 				if tPunc.Value == l.PUNCTUATION_LPAREN {
// 					depth++
// 				} else if tPunc.Value == l.PUNCTUATION_RPAREN {
// 					if depth > 0 {
// 						depth--
// 					} else {
// 						return NodeBinary{}, ErrMismatchedParenthesis
// 					}
// 				}
// 			}
// 		}
// 	}

// 	if !isLeftRead {
// 		return NodeBinary{}, fmt.Errorf("not a binary expression")
// 	}

// 	leftParsed, err := parseExpression(left)
// 	if err != nil {
// 		return NodeBinary{}, err
// 	}

// 	rightParsed, err := parseExpression(right)
// 	if err != nil {
// 		return NodeBinary{}, err
// 	}

// 	return NodeBinary{
// 		Operator: oper.Value,
// 		Left:     leftParsed,
// 		Right:    rightParsed,
// 	}, nil
// }

// func parseAssignExpression(expr []scanner.Token) (NodeAssign, error) {
// 	node := NodeAssign{}

// 	if len(expr) < 3 {
// 		return node, fmt.Errorf("not enough tokens to parse expression as assign")
// 	}

// 	// start

// 	v, ok := expr[0].(scanner.TokenWord)
// 	if !ok {
// 		return node, fmt.Errorf("expected a variable")
// 	}
// 	if l.IsConstant(v.Value) {
// 		return node, ErrVarAsConst
// 	}

// 	node.Left = NodeVariable{
// 		Name: v.Value,
// 	}

// 	// assign operator

// 	oper, ok := expr[1].(scanner.TokenOperator)
// 	if !ok || oper.Value != l.OPERATOR_ASS {
// 		return node, fmt.Errorf("expected an assign operator")
// 	}

// 	// rest

// 	right, err := parseExpression(expr[2:])
// 	if err != nil {
// 		return node, err
// 	}

// 	node.Right = right

// 	return node, nil
// }

// func isIRangeFunctionArg(arg []scanner.Token) bool {
// 	if len(arg) >= 3 {
// 		_, ok := arg[0].(scanner.TokenWord)
// 		if !ok {
// 			return false
// 		}

// 		argAss, ok := arg[1].(scanner.TokenOperator)
// 		if !ok || argAss.Value != l.OPERATOR_ASS {
// 			return false
// 		}

// 		_, err := parseRange(arg[2:])
// 		if err != nil {
// 			return false
// 		}

// 		return true
// 	}

// 	return false
// }

// func parseFunctionArgs(tl []scanner.Token) ([]Node, error) {
// 	args := []Node{}

// 	sliced, err := sliceTokenListByComma(tl)
// 	if err != nil {
// 		return args, err
// 	}

// 	for _, arg := range sliced {
// 		if len(arg) > 0 {
// 			var parsedArg Node

// 			if isIRangeFunctionArg(arg) {
// 				parsedArg, err = parseAssignExpression(arg)
// 				if err != nil {
// 					return args, err
// 				}
// 			} else {
// 				parsedArg, err = parseExpression(parenthesiseExpression(arg))
// 				if err != nil {
// 					return args, err
// 				}
// 			}

// 			args = append(args, parsedArg)
// 		}
// 	}

// 	return args, nil
// }

// // TODO: switch from full parenthesisation to this -> https://en.wikipedia.org/wiki/Shunting_yard_algorithm

// func parseExpression(expr []scanner.Token) (Node, error) {
// 	if len(expr) == 0 {
// 		return InvalidNode{}, fmt.Errorf("empty expression")
// 	}

// 	switch firstToken := expr[0].(type) {
// 	case scanner.TokenPunctuation:
// 		if firstToken.Value == l.PUNCTUATION_LPAREN {
// 			if isBinary(expr) {
// 				return parseBinary(expr)
// 			} else {
// 				return parseExpression(readParentheses(expr))
// 			}
// 		}

// 	case scanner.TokenNumber:
// 		return NodeNumber{
// 			Value: firstToken.Value,
// 		}, nil

// 	case scanner.TokenWord:
// 		if len(expr) > 1 {
// 			switch followUpToken := expr[1].(type) {
// 			case scanner.TokenOperator:
// 				{
// 					switch followUpToken.Value {
// 					case l.OPERATOR_ASS:
// 						// this is an assign expression
// 						return parseAssignExpression(expr)

// 					default:
// 						return InvalidNode{}, fmt.Errorf("word followed by operator token with unexpected value %s", followUpToken.Value)
// 					}
// 				}

// 			case scanner.TokenPunctuation:
// 				{
// 					switch followUpToken.Value {
// 					case l.PUNCTUATION_LPAREN:
// 						// this is a function
// 						args, err := parseFunctionArgs(readParentheses(expr))
// 						return NodeFuncCall{
// 							Name:      firstToken.Value,
// 							Arguments: args,
// 						}, err

// 					default:
// 						return InvalidNode{}, fmt.Errorf("word followed by punctuation token with unexpected value %s", followUpToken.Value)
// 					}
// 				}

// 			default:
// 				return InvalidNode{}, fmt.Errorf("word followed by unexpected token %s", followUpToken.Type())
// 			}

// 		} else {
// 			if l.IsConstant(firstToken.Value) {
// 				return NodeConstant{
// 					Name: firstToken.Value,
// 				}, nil
// 			} else {
// 				return NodeVariable{
// 					Name: firstToken.Value,
// 				}, nil
// 			}
// 		}
// 	}

// 	return InvalidNode{}, fmt.Errorf("failed to parse an expression")
// }

// func Parse(tl []scanner.Token) (NodeRoot, error) {
// 	normalised := normalise(tl)
// 	v, err := parseExpression(parenthesiseExpression(normalised))

// 	return NodeRoot{
// 		Value: v,
// 	}, err
// }
