package interpreter

import "math"

const (
	psi = 1.46557123187676802665 // https://oeis.org/A092526
	//    1.465571231876768026656731
)

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

func IsConstant(name string) (ConstantDefinition, bool) {
	for _, def := range DEFINED_CONSTANTS {
		if def.Name == name {
			return def, true
		}
	}

	return ConstantDefinition{}, false
}
