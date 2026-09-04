package interpreter

import (
	"fmt"
	"gocalc/interpreter/lexemes"
	"gocalc/interpreter/parser"
	"gocalc/interpreter/scanner"
	"math"
)

func solveBinary(node parser.NodeBinary) (float64, error) {
	left, err := solveNode(node.Left)
	if err != nil {
		return 0, err
	}

	right, err := solveNode(node.Right)
	if err != nil {
		return 0, err
	}

	switch node.Operator {
	case lexemes.OPERATOR_ADD:
		return left + right, nil

	case lexemes.OPERATOR_SUB:
		return left - right, nil

	case lexemes.OPERATOR_MUL:
		return left * right, nil

	case lexemes.OPERATOR_DIV:
		return left / right, nil

	case lexemes.OPERATOR_POW:
		return math.Pow(left, right), nil

	default:
		return 0, fmt.Errorf("undefined operator \"%s\"", node.Operator)
	}
}

func solveNode(node parser.Node) (float64, error) {
	nodeBin, isBin := node.(parser.NodeBinary)
	if isBin {
		return solveBinary(nodeBin)
	}

	switch typedNode := node.(type) {
	case parser.NodeNumber:
		return typedNode.Value, nil

	case parser.NodeConstant:
		{
			constName := typedNode.Name

			switch constName {
			case lexemes.CONSTANT_PI:
				return math.Pi, nil

			case lexemes.CONSTANT_E:
				return math.E, nil

			case lexemes.CONSTANT_PHI:
				return math.Phi, nil

			default:
				return 0, fmt.Errorf("undefined constant \"%s\"", constName)
			}
		}

	case parser.NodeFuncCall:
		{
			funcName := typedNode.Name

			_, isIRangeFunc := typedNode.Arguments[0].(parser.NodeIRangeFuncMainArg)
			if isIRangeFunc {
				//
			}

			funcArgs := []float64{}
			for _, arg := range typedNode.Arguments {
				f, err := solveNode(arg)
				if err != nil {
					return 0, err
				}

				funcArgs = append(funcArgs, f)
			}

			argc := len(funcArgs)

			switch funcName {
			case lexemes.FUNCTION_SIN:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Sin(funcArgs[0]), nil

			case lexemes.FUNCTION_COS:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Cos(funcArgs[0]), nil

			case lexemes.FUNCTION_TAN:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Tan(funcArgs[0]), nil

			case lexemes.FUNCTION_ATAN:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Atan(funcArgs[0]), nil

			case lexemes.FUNCTION_EXP:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Exp(funcArgs[0]), nil

			case lexemes.FUNCTION_ABS:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Abs(funcArgs[0]), nil

			case lexemes.FUNCTION_LOG:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Log10(funcArgs[0]), nil

			case lexemes.FUNCTION_LN:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Log(funcArgs[0]), nil

			case lexemes.FUNCTION_SQRT:
				if argc != 1 {
					return 0, fmt.Errorf("%s function expected 1 argument", funcName)
				}
				return math.Sqrt(funcArgs[0]), nil

			default:
				return 0, fmt.Errorf("undefined function \"%s\"", funcName)
			}
		}

	default:
		return 0, fmt.Errorf("undefined node type \"%s\"", node.Type())
	}
}

func CompileToAST(s string) (parser.NodeRoot, error) {
	tl, err := scanner.Scan(s)
	if err != nil {
		return parser.NodeRoot{}, err
	}

	ast, err := parser.Parse(tl)
	if err != nil {
		return parser.NodeRoot{}, err
	}

	return ast, nil
}

func EvaluateAST(ast parser.NodeRoot) (float64, error) {
	return solveNode(ast.Value)
}

func EvaluateString(s string) (float64, error) {
	ast, err := CompileToAST(s)
	if err != nil {
		return 0, err
	}

	return EvaluateAST(ast)
}
