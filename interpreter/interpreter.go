package interpreter

import (
	"fmt"
	"gocalc/interpreter/parser"
	"gocalc/interpreter/scanner"
	"math"
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
		case scanner.ALLOWED_CONSTANTS[PI]:
			return math.Pi, nil

		case scanner.ALLOWED_CONSTANTS[E]:
			return math.E, nil

		case scanner.ALLOWED_CONSTANTS[PHI]:
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
		case scanner.ALLOWED_FUNCTIONS[SIN]:
			return math.Sin(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[COS]:
			return math.Cos(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[TAN]:
			return math.Tan(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[ATAN]:
			return math.Atan(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[EXP]:
			return math.Exp(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[ABS]:
			return math.Abs(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[LOG]:
			return math.Log10(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[LN]:
			return math.Log(funcArg), nil

		case scanner.ALLOWED_FUNCTIONS[SQRT]:
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

	const (
		ADD = iota
		SUB
		MUL
		DIV
		POW
	)

	switch binOper {
	case scanner.ALLOWED_OPERATORS[ADD]:
		return binLeft + binRight, nil

	case scanner.ALLOWED_OPERATORS[SUB]:
		return binLeft - binRight, nil

	case scanner.ALLOWED_OPERATORS[MUL]:
		return binLeft * binRight, nil

	case scanner.ALLOWED_OPERATORS[DIV]:
		return binLeft / binRight, nil

	case scanner.ALLOWED_OPERATORS[POW]:
		return math.Pow(binLeft, binRight), nil

	default:
		return 0, fmt.Errorf("Undefined operator \"%s\".", string(binOper))
	}
}

func CompileToAST(s string) (parser.Node, error) {
	sc, err := scanner.Scan(s)
	if err != nil {
		return parser.Node{}, err
	}

	p, err := parser.Parse(sc)
	if err != nil {
		return parser.Node{}, err
	}

	return p, nil
}

func EvaluateAST(ast parser.Node) (float64, error) {
	if ast.Type != parser.NODE_TYPE_ROOT {
		return 0, fmt.Errorf("Expected an AST root, but got \"%s\".", ast.Type)
	}

	v, ok := ast.Value.(parser.Node)
	if !ok {
		return 0, fmt.Errorf("Failed type assertion : ASTRoot.Value is not a parser.Node.")
	}

	return solveNode(v)
}

func EvaluateString(s string) (float64, error) {
	ast, err := CompileToAST(s)
	if err != nil {
		return 0, err
	}

	return EvaluateAST(ast)
}
