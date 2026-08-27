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
		{
			n, ok := node.Value.(float64)
			if ok {
				return n, nil
			} else {
				return 0, fmt.Errorf("Failed type assertion : number node value is not a float64.")
			}
		}

	case parser.NODE_TYPE_CONSTANT:
		{
			nodeValue := node.Value.(parser.NodeValueConstant)

			constName := nodeValue.Name

			switch constName {
			case scanner.CONSTANT_PI:
				return math.Pi, nil

			case scanner.CONSTANT_E:
				return math.E, nil

			case scanner.CONSTANT_PHI:
				return math.Phi, nil

			default:
				return 0, fmt.Errorf("Undefined constant \"%s\".", constName)
			}
		}

	case parser.NODE_TYPE_FUNCTION_CALL:
		{
			nodeValue := node.Value.(parser.NodeValueFunction)

			funcName := nodeValue.Name
			funcArg, err := solveNode(nodeValue.Argument)
			if err != nil {
				return 0, err
			}

			switch funcName {
			case scanner.FUNCTION_SIN:
				return math.Sin(funcArg), nil

			case scanner.FUNCTION_COS:
				return math.Cos(funcArg), nil

			case scanner.FUNCTION_TAN:
				return math.Tan(funcArg), nil

			case scanner.FUNCTION_ATAN:
				return math.Atan(funcArg), nil

			case scanner.FUNCTION_EXP:
				return math.Exp(funcArg), nil

			case scanner.FUNCTION_ABS:
				return math.Abs(funcArg), nil

			case scanner.FUNCTION_LOG:
				return math.Log10(funcArg), nil

			case scanner.FUNCTION_LN:
				return math.Log(funcArg), nil

			case scanner.FUNCTION_SQRT:
				return math.Sqrt(funcArg), nil

			default:
				return 0, fmt.Errorf("Undefined function \"%s\".", funcName)
			}
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
	case scanner.OPERATOR_ADD:
		return binLeft + binRight, nil

	case scanner.OPERATOR_SUB:
		return binLeft - binRight, nil

	case scanner.OPERATOR_MUL:
		return binLeft * binRight, nil

	case scanner.OPERATOR_DIV:
		return binLeft / binRight, nil

	case scanner.OPERATOR_POW:
		return math.Pow(binLeft, binRight), nil

	default:
		return 0, fmt.Errorf("Undefined operator \"%s\".", binOper)
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
		return 0, fmt.Errorf("Failed type assertion : ASTRoot.Value is not a node.")
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
