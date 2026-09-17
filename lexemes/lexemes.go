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

// Punctuation //------------------------------------------------------------//
