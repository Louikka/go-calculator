package main

import (
	"fmt"
	"gocalc/lexer"
	"gocalc/parser"
)

func main() {
	l, err := lexer.Analyse("123 + 4 * COS(8)")
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
