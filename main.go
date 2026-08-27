package main

import (
	"bufio"
	"flag"
	"fmt"
	"gocalc/interpreter"
	"gocalc/lib"
	"os"
	"slices"
	"strings"
)

type ProgramFlags struct {
	loop bool
	ast  bool
	// File name to write program output
	out string
}

func initPropgramFlags() *ProgramFlags {
	r := flag.Bool("r", false, "repeat input")
	ast := flag.Bool("ast", false, "output AST instead of calculating result")
	out := flag.String("o", "out.txt", "specify output file")

	flag.Parse()

	return &ProgramFlags{
		loop: *r,
		ast:  *ast,
		out:  *out,
	}
}

func checkIfStop(s string) bool {
	stop := []string{"Q", "QUIT", "STOP", "END"}
	s_prep := strings.ToUpper(strings.TrimSpace(s))

	return slices.Contains(stop, s_prep)
}

func main() {
	flags := initPropgramFlags()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while trying to read input :", err)
			return
		}

		if checkIfStop(s) {
			break
		}

		ast, err := interpreter.CompileToAST(s)
		if err != nil {
			fmt.Println(err)
			return
		}

		if flags.ast {
			err = lib.WriteJSONToFile(ast, flags.out)
			if err != nil {
				fmt.Println("Failed to write parsed AST to file :", err)
			} else {
				fmt.Printf("Parsed AST written to %s\n", flags.out)
			}

			return
		}

		n, err := interpreter.EvaluateAST(ast)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(n)

		if !flags.loop {
			break
		}
	}
}
