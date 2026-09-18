package scanner

import "slices"

// Operators //--------------------------------------------------------------//

const (
	ASSOC_UNDEFINED = ""
	ASSOC_LEFT      = "LEFT"
	ASSOC_RIGHT     = "RIGHT"
)

type OperatorDefinition struct {
	Value         string
	Precedence    int
	Associativity string
}

func (d OperatorDefinition) Len() int {
	return len(d.Value)
}

var DEFINED_OPERATORS = []OperatorDefinition{
	{
		Value:         "+",
		Precedence:    1,
		Associativity: ASSOC_LEFT,
	},
	{
		Value:         "-",
		Precedence:    1,
		Associativity: ASSOC_LEFT,
	},
	{
		Value:         "*",
		Precedence:    2,
		Associativity: ASSOC_LEFT,
	},
	{
		Value:         "/",
		Precedence:    2,
		Associativity: ASSOC_LEFT,
	},
	{
		Value:         "^",
		Precedence:    3,
		Associativity: ASSOC_RIGHT,
	},
	{
		Value:         "=",
		Precedence:    0,
		Associativity: ASSOC_LEFT,
	},
	{
		Value:         "..",
		Precedence:    999,
		Associativity: ASSOC_LEFT,
	},
}

func IsOperator(s string) (OperatorDefinition, bool) {
	for _, def := range DEFINED_OPERATORS {
		if def.Value == s {
			return def, true
		}
	}

	return OperatorDefinition{}, false
}

// Length of the longest operator (in bytes).
var LongestOperatorLen = LongestLen(DEFINED_OPERATORS)

var OperatorsPrecedenceSorted = func() []int {
	l := []int{}

	for _, oper := range DEFINED_OPERATORS {
		if !slices.Contains(l, oper.Precedence) {
			l = append(l, oper.Precedence)
		}
	}

	slices.Sort(l)

	return l
}()

// Punctuation //------------------------------------------------------------//

type PunctuationDefinition struct {
	Value string
}

func (d PunctuationDefinition) Len() int {
	return len(d.Value)
}

var DEFINED_PUCTUATION = []PunctuationDefinition{
	{
		Value: "(",
	},
	{
		Value: ")",
	},
	{
		Value: ",",
	},
}

func IsPunctuation(s string) (PunctuationDefinition, bool) {
	for _, def := range DEFINED_PUCTUATION {
		if def.Value == s {
			return def, true
		}
	}

	return PunctuationDefinition{}, false
}

// Length of the longest punctuation (in bytes).
var LongestPunctuationLen = LongestLen(DEFINED_PUCTUATION)
