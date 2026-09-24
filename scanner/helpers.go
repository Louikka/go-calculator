package scanner

import "strings"

type HaveLenMethod interface {
	Len() int
}

func LongestLen[T HaveLenMethod](ol []T) int {
	l := 0

	for _, o := range ol {
		new := o.Len()
		if new > l {
			l = new
		}
	}

	return l
}

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
