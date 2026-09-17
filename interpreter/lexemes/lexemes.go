package lexemes

import (
	"math"
)

const (
	psi = 1.46557123187676802665 // https://oeis.org/A092526
	//    1.465571231876768026656731
)

// Constants //--------------------------------------------------------------//

type ConstantDefinition struct {
	Name  string
	Value float64
}

var DEFINED_CONSTANTS = []ConstantDefinition{
	{
		Name:  "PI",
		Value: math.Pi,
	},
	{
		Name:  "E",
		Value: math.E,
	},
	{
		Name:  "PHI",
		Value: math.Phi,
	},
	{
		Name:  "PSI",
		Value: psi,
	},
}

// This function checks if constant s is defined.
func IsConstant(s string) (ConstantDefinition, bool) {
	for _, def := range DEFINED_CONSTANTS {
		if def.Name == s {
			return def, true
		}
	}

	return ConstantDefinition{}, false
}

// Functions //--------------------------------------------------------------//

type FunctionDefinition struct {
	Name string
	// Arguments count (how many arguments function takes).
	Argc int
}

var DEFINED_FUNCTIONS = []FunctionDefinition{
	{
		Name: "SIN",
		Argc: 1,
	},
	{
		Name: "COS",
		Argc: 1,
	},
	{
		Name: "TAN",
		Argc: 1,
	},
	{
		Name: "ATAN",
		Argc: 1,
	},
	{
		Name: "ABS",
		Argc: 1,
	},
	{
		Name: "LOG",
		Argc: 1,
	},
	{
		Name: "LN",
		Argc: 1,
	},
	{
		Name: "SQRT",
		Argc: 1,
	},
	{
		Name: "CBRT",
		Argc: 1,
	},
	{
		Name: "ROUND",
		Argc: 1,
	},
	{
		Name: "RAND",
		Argc: 0,
	},
	{
		Name: "SUM",
		Argc: 2,
	},
	{
		Name: "PROD",
		Argc: 2,
	},
}

func IsFunction(s string) (FunctionDefinition, bool) {
	for _, def := range DEFINED_FUNCTIONS {
		if def.Name == s {
			return def, true
		}
	}

	return FunctionDefinition{}, false
}

// Operators //--------------------------------------------------------------//

type OperatorDefinition struct {
	Value             string
	Precedence        int
	IsLeftAssociative bool
}

var DEFINED_OPERATORS = []OperatorDefinition{
	{
		Value:             "+",
		Precedence:        1,
		IsLeftAssociative: true,
	},
	{
		Value:             "-",
		Precedence:        1,
		IsLeftAssociative: true,
	},
	{
		Value:             "*",
		Precedence:        2,
		IsLeftAssociative: true,
	},
	{
		Value:             "/",
		Precedence:        2,
		IsLeftAssociative: true,
	},
	{
		Value:             "^",
		Precedence:        3,
		IsLeftAssociative: false,
	},
	{
		Value:             "=",
		Precedence:        0,
		IsLeftAssociative: true,
	},
	{
		Value:             "..",
		Precedence:        999,
		IsLeftAssociative: true,
	},
}

// Length of the longest operator (in bytes).
var LONGEST_OPERATOR_LEN = func() int {
	length := 0

	for _, oper := range DEFINED_OPERATORS {
		newLength := len(oper.Value)
		if newLength > length {
			length = newLength
		}
	}

	return length
}()

func IsOperator(s string) (OperatorDefinition, bool) {
	for _, def := range DEFINED_OPERATORS {
		if def.Value == s {
			return def, true
		}
	}

	return OperatorDefinition{}, false
}

// Punctuation //------------------------------------------------------------//

type PunctuationDefinition struct {
	Value string
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

// Length of the longest punctuation (in bytes).
var LONGEST_PUCTUATION_LEN = func() int {
	length := 0

	for _, punc := range DEFINED_PUCTUATION {
		newLength := len(punc.Value)
		if newLength > length {
			length = newLength
		}
	}

	return length
}()

func IsPunctuation(s string) (PunctuationDefinition, bool) {
	for _, def := range DEFINED_PUCTUATION {
		if def.Value == s {
			return def, true
		}
	}

	return PunctuationDefinition{}, false
}
