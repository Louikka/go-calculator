package main

import (
	"bufio"
	"flag"
	"fmt"
	"gocalc/interpreter"
	"os"
	"slices"
	"strings"
)

type ProgramFlags struct {
	// If program should loop input prompt until user stop or error.
	loop bool
}

func initPropgramFlags() ProgramFlags {
	r := flag.Bool("r", false, "repeat input prompt")

	flag.Parse()

	return ProgramFlags{
		loop: *r,
	}
}

func checkIfStop(s string) bool {
	stopKW := []string{"Q", "QUIT", "STOP", "END"}
	inp := strings.ToUpper(strings.TrimSpace(s))
	return slices.Contains(stopKW, inp)
}

func main() {
	flags := initPropgramFlags()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while trying to read input:", err)
			return
		}

		if checkIfStop(s) {
			break
		}

		n, err := interpreter.EvaluateString(s)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(n)

		if !flags.loop {
			break
		}
	}
}
