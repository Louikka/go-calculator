package lexemes

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
var LONGEST_OPERATOR_LEN = func() int {
	maxLen := 0

	for _, oper := range DEFINED_OPERATORS {
		newLen := oper.Len()
		if newLen > maxLen {
			maxLen = newLen
		}
	}

	return maxLen
}()
