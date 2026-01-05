package parser

import (
	"errors"
	"gocalc/lexer"
)

/* Types and constants *******************************************************/

const (
	TOKEN_TYPE_NUMBER      = lexer.TOKEN_TYPE_NUMBER
	TOKEN_TYPE_OPERATOR    = lexer.TOKEN_TYPE_OPERATOR
	TOKEN_TYPE_KEYWORD     = lexer.TOKEN_TYPE_KEYWORD
	TOKEN_TYPE_PUNCTUATION = lexer.TOKEN_TYPE_PUNCTUATION
)

type NodeType string

const (
	NODE_TYPE_ROOT     NodeType = "ROOT"
	NODE_TYPE_NUMBER   NodeType = "NUMBER"
	NODE_TYPE_FUNCTION NodeType = "FUNCTION"
	NODE_TYPE_CONSTANT NodeType = "CONSTANT"
	NODE_TYPE_BINARY   NodeType = "BINARY"
)

type Node struct {
	Type  NodeType
	Value any
}

type NodeValueNumber struct {
	Value float64
}
type NodeValueFunction struct {
	Name     string
	Argument Node
}
type NodeValueConstant struct {
	Name string
}
type NodeValueBinary struct {
	Operator byte
	Left     Node
	Right    Node
}

/*
type NumberNode struct {
	Type  NodeType
	Value float64
}
type FunctionNode struct {
	Type     NodeType
	Name     string
	Argument Node
}
type ConstantNode struct {
	Type NodeType
	Name string
}
type BinaryNode struct {
	Type     NodeType
	Operator byte
	Left     Node
	Right    Node
}
*/

/* Custom token list reader **************************************************/

type _TokenListReader struct {
	tl  []lexer.Token
	pos uint
}

func (r *_TokenListReader) peek(offset int) (lexer.Token, error) {
	newPos := int(r.pos) + offset

	if newPos < 0 || newPos >= len(r.tl) {
		return lexer.Token{}, errors.New("Position is out of bounds.")
	}

	return r.tl[newPos], nil
}

func (r *_TokenListReader) next() (lexer.Token, error) {
	r.pos++

	if r.pos >= uint(len(r.tl)) {
		return lexer.Token{}, errors.New("Position is out of bounds.")
	}

	return r.tl[r.pos], nil
}

func (r *_TokenListReader) isEndOfString() bool {
	return r.pos >= uint(len(r.tl))
}

/* Helpers *******************************************************************/

func repeatInSlice[T any](item T, count uint) []T {
	// Pre-allocate the slice with a specific length
	s := make([]T, count)

	for i := range s {
		s[i] = item
	}

	return s
}

func parenthesizeExpression(input []lexer.Token) []lexer.Token {
	parenOpen := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('('),
	}
	parenClose := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte(')'),
	}
	caret := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('^'),
	}
	star := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('*'),
	}
	slash := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('/'),
	}
	plus := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('+'),
	}
	minus := lexer.Token{
		Type:  TOKEN_TYPE_PUNCTUATION,
		Value: byte('-'),
	}

	output := []lexer.Token{}

	output = append(output, repeatInSlice(parenOpen, 4)...)

	for i, v := range input {
		t := v

		if t.Type == TOKEN_TYPE_OPERATOR || t.Type == TOKEN_TYPE_PUNCTUATION {
			switch t.Value.(byte) {
			case '(':
				output = append(output, repeatInSlice(parenOpen, 4)...)
				continue
			case ')':
				output = append(output, repeatInSlice(parenClose, 4)...)
				continue
			case '^':
				output = append(output, parenClose, caret, parenOpen)
				continue
			case '*':
				output = append(output, parenClose, parenClose, star, parenOpen, parenOpen)
				continue
			case '/':
				output = append(output, parenClose, parenClose, slash, parenOpen, parenOpen)
				continue
			case '+':
				//fmt.Println(i, v)
				// unary check: either first or had an operator expecting secondary argument
				if i == 0 || input[i-1].Type == TOKEN_TYPE_PUNCTUATION {
					output = append(output, plus)
				} else {
					output = append(output, parenClose, parenClose, parenClose, plus, parenOpen, parenOpen, parenOpen)
				}

				continue
			case '-':
				if i == 0 || input[i-1].Type == TOKEN_TYPE_PUNCTUATION {
					output = append(output, minus)
				} else {
					output = append(output, parenClose, parenClose, parenClose, minus, parenOpen, parenOpen, parenOpen)
				}

				continue
			}
		}

		output = append(output, t)
	}

	output = append(output, repeatInSlice(parenClose, 4)...)

	return output
}

/* Parser main ***************************************************************/

func Parse(tl []lexer.Token) (Node, error) {
	//fmt.Println(lexer.StringifyTokens(parenthesizeExpression(tl)))

	return Node{
		Type:  NODE_TYPE_ROOT,
		Value: nil,
	}, nil
}
