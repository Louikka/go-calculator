package scanner

import "strings"

func Stringify(tl []Token, delimeter string) string {
	var s strings.Builder

	last := len(tl) - 1

	for i, t := range tl {
		a := t.String()

		if i < last {
			a += delimeter
		}

		s.WriteString(a)
	}

	return s.String()
}
