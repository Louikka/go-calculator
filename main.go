package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"gocalc/interpreter"
	"gocalc/parser"
	"gocalc/scanner"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type ProgramFlags struct {
	loop           bool
	ast            bool
	outputFilename string
}

func initPropgramFlags() *ProgramFlags {
	r := flag.Bool("r", false, "repeat input")
	ast := flag.Bool("ast", false, "output AST instead of calculating result")
	out := flag.String("o", "out.txt", "specify output file")

	flag.Parse()

	return &ProgramFlags{
		loop:           *r,
		ast:            *ast,
		outputFilename: *out,
	}
}

func checkIfStop(s string) bool {
	stop := []string{"Q", "QUIT", "STOP", "END"}
	s_prep := strings.ToUpper(strings.TrimSpace(s))

	return slices.Contains(stop, s_prep)
}

func outputASTInFile(s string, filename string) error {
	tl, err := scanner.Scan(s)
	if err != nil {
		return err
	}

	ast, err := parser.Parse(tl)
	if err != nil {
		return err
	}

	ast_b, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, ast_b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flags := initPropgramFlags()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error : An error occured while trying to read input :", err)
			return
		}

		if flags.ast {
			err := outputASTInFile(s, flags.outputFilename)
			if err != nil {
				fmt.Println("Error : Failed to write AST in file :", err)
			} else {
				fmt.Printf("Parsed AST written to %s", flags.outputFilename)
			}

			return
		}

		if checkIfStop(s) {
			break
		}

		n, err := interpreter.Calculate(s)
		if err != nil {
			fmt.Println("Error :", err)
			return
		}

		fmt.Println(n)

		if !flags.loop {
			break
		}
	}
}
