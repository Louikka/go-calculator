package main

import (
	"bufio"
	"fmt"
	"gocalc/lexer"
	"gocalc/parser"
	"math"
	"os"
)

func solveNode(node parser.Node) (float64, error) {

	if node.Type == parser.NODE_TYPE_BINARY {
		return solveBinary(node)
	}

	switch node.Type {

	case parser.NODE_TYPE_NUMBER:
		//nodeValue := node.Value.(parser.NodeValueNumber)
		return node.Value.(float64), nil

	case parser.NODE_TYPE_CONSTANT:
		nodeValue := node.Value.(parser.NodeValueConstant)

		constName := nodeValue.Name

		const (
			PI = iota
			E
			PHI
		)

		switch constName {
		case lexer.ALLOWED_CONSTANTS[PI]:
			return math.Pi, nil
		case lexer.ALLOWED_CONSTANTS[E]:
			return math.E, nil
		case lexer.ALLOWED_CONSTANTS[PHI]:
			return math.Phi, nil
		default:
			return 0, fmt.Errorf("Undefined constant \"%s\".", constName)
		}

	case parser.NODE_TYPE_FUNCTION:
		nodeValue := node.Value.(parser.NodeValueFunction)

		funcName := nodeValue.Name
		funcArg, err := solveNode(nodeValue.Argument)
		if err != nil {
			return 0, err
		}

		const (
			SIN = iota
			COS
			TAN
			ATAN
			EXP
			ABS
			LOG
			LN
			SQRT
		)

		switch funcName {
		case lexer.ALLOWED_FUNCTIONS[SIN]:
			return math.Sin(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[COS]:
			return math.Cos(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[TAN]:
			return math.Tan(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[ATAN]:
			return math.Atan(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[EXP]:
			return math.Exp(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[ABS]:
			return math.Abs(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[LOG]:
			return math.Log10(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[LN]:
			return math.Log(funcArg), nil
		case lexer.ALLOWED_FUNCTIONS[SQRT]:
			return math.Sqrt(funcArg), nil
		default:
			return 0, fmt.Errorf("Undefined function \"%s\".", funcName)
		}

	default:
		return 0, fmt.Errorf("Undefined node type \"%s\".", node.Type)

	}
}

func solveBinary(node parser.Node) (float64, error) {
	if node.Type != parser.NODE_TYPE_BINARY {
		return 0, fmt.Errorf("The node type is \"%s\", expected \"%s\".", node.Type, parser.NODE_TYPE_BINARY)
	}

	nodeValue := node.Value.(parser.NodeValueBinary)

	binOper := nodeValue.Operator
	binLeft, err := solveNode(nodeValue.Left)
	if err != nil {
		return 0, err
	}
	binRight, err := solveNode(nodeValue.Right)
	if err != nil {
		return 0, err
	}

	switch binOper {

	case lexer.ALLOWED_OPERATORS[0]: // +
		return binLeft + binRight, nil

	case lexer.ALLOWED_OPERATORS[1]: // -
		return binLeft - binRight, nil

	case lexer.ALLOWED_OPERATORS[2]: // *
		return binLeft * binRight, nil

	case lexer.ALLOWED_OPERATORS[3]: // /
		return binLeft / binRight, nil

	case lexer.ALLOWED_OPERATORS[4]: // ^
		return math.Pow(binLeft, binRight), nil

	default:
		return 0, fmt.Errorf("Undefined operator \"%s\".", string(binOper))

	}
}

func calculate(s string) (float64, error) {
	l, err := lexer.Analyse(s)
	if err != nil {
		return 0, err
	}

	p, err := parser.Parse(l)
	if err != nil {
		return 0, err
	}

	if p.Type != parser.NODE_TYPE_ROOT {
		return 0, fmt.Errorf("Expected a root of the ASTree.")
	}

	return solveNode(p.Value.(parser.Node))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">>> ")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while trying to read input : ", err)
		return
	}

	n, err := calculate(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}
