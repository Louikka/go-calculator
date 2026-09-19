package interpreter

import (
	"math"
	"math/rand/v2"
)

// functions without arguments //--------------------------------------------//

type FunctionDefinition0 struct {
	Name string
	Fn   func() float64
}

func (d FunctionDefinition0) Argc() uint {
	return 0
}

var DEFINED_FUNCTIONS_WITHOUT_ARGUMENTS = []FunctionDefinition0{
	{
		Name: "RAND",
		Fn:   rand.Float64,
	},
}

func IsFunc0(name string) (FunctionDefinition0, bool) {
	for _, def := range DEFINED_FUNCTIONS_WITHOUT_ARGUMENTS {
		if def.Name == name {
			return def, true
		}
	}

	return FunctionDefinition0{}, false
}

// functions with 1 argument //----------------------------------------------//

type FunctionDefinition1 struct {
	Name string
	Fn   func(n float64) float64
}

func (d FunctionDefinition1) Argc() uint {
	return 1
}

var DEFINED_FUNCTIONS_WITH_1_ARGUMENT = []FunctionDefinition1{
	{
		Name: "SIN",
		Fn:   math.Sin,
	},
	{
		Name: "COS",
		Fn:   math.Cos,
	},
	{
		Name: "TAN",
		Fn:   math.Tan,
	},
	{
		Name: "ATAN",
		Fn:   math.Atan,
	},
	{
		Name: "ABS",
		Fn:   math.Abs,
	},
	{
		Name: "LOG",
		Fn:   math.Log10,
	},
	{
		Name: "LN",
		Fn:   math.Log,
	},
	{
		Name: "SQRT",
		Fn:   math.Sqrt,
	},
	{
		Name: "CBRT",
		Fn:   math.Cbrt,
	},
	{
		Name: "ROUND",
		Fn:   math.Round,
	},
}

func IsFunc1(name string) (FunctionDefinition1, bool) {
	for _, def := range DEFINED_FUNCTIONS_WITH_1_ARGUMENT {
		if def.Name == name {
			return def, true
		}
	}

	return FunctionDefinition1{}, false
}
