package main

import (
	"fmt"
	"gocalc/lexer"
	"gocalc/parser"
)

func main() {
	const s = "123 + 4 * COS(8)"
	fmt.Println(s)

	l, err := lexer.Analyse(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l)

	p, err := parser.Parse(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
