package main

import (
	"fmt"
	"gocalc/lexer"
)

func main() {
	l, err := lexer.Analyse("123 + 456 * SQRT(4)")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l)
}
