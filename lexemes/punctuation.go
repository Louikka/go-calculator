package lexemes

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
var LONGEST_PUCTUATION_LEN = func() int {
	length := 0

	for _, punc := range DEFINED_PUCTUATION {
		newLength := punc.Len()
		if newLength > length {
			length = newLength
		}
	}

	return length
}()
