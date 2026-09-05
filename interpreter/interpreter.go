package interpreter

import (
	"fmt"
	"gocalc/interpreter/lexemes"
	"gocalc/interpreter/parser"
	"gocalc/interpreter/scanner"
	"math"
)

var psi float64 = (1 + math.Cbrt(29+3*math.Sqrt(93)) + math.Cbrt(29-3*math.Sqrt(93))/2.0) / 3.0

func solveNodeConstant(node parser.NodeConstant) (float64, error) {
	switch node.Name {
	case lexemes.CONSTANT_PI:
		return math.Pi, nil

	case lexemes.CONSTANT_E:
		return math.E, nil

	case lexemes.CONSTANT_PHI:
		return math.Phi, nil

	case lexemes.CONSTANT_PSI:
		return 1.46557123187676802665, nil

	default:
		return 0, fmt.Errorf("undefined constant \"%s\"", node.Name)
	}
}

type _VariableContext struct {
	Name  string
	Value float64
}

func solveNodeDefaultFuncCall(node parser.NodeFuncCall) (float64, error) {
	funcName := node.Name
	funcArgc := node.Argc()
	funcArgs := []float64{}

	for _, arg := range node.Arguments {
		f, err := solveNode(arg, []_VariableContext{})
		if err != nil {
			return 0, err
		}

		funcArgs = append(funcArgs, f)
	}

	switch funcName {
	case lexemes.FUNCTION_SIN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Sin(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_COS:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Cos(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_TAN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Tan(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_ATAN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Atan(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_EXP:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Exp(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_ABS:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Abs(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_LOG:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Log10(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_LN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Log(funcArgs[0]), nil
		}

	case lexemes.FUNCTION_SQRT:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Sqrt(funcArgs[0]), nil
		}

	default:
		return 0, fmt.Errorf("undefined function \"%s\"", funcName)
	}
}

func solveNodeIRangeFuncCall(node parser.NodeFuncCall) (float64, error) {
	funcName := node.Name
	funcArgc := node.Argc()

	if funcArgc != 2 {
		return 0, fmt.Errorf("%s expected 2 arguments, but got %d", funcName, funcArgc)
	}

	mainArg, isIRangeFunc := node.Arguments[0].(parser.NodeIRangeFuncMainArg)
	if !isIRangeFunc {
		return 0, fmt.Errorf("%s expected variable and range as first argument, but got %s", funcName, node.Arguments[0].Type())
	}

	secondArg := node.Arguments[1]

	switch funcName {
	case lexemes.FUNCTION_SUM:
		{
			var sum float64 = 0
			for i := mainArg.Range.Start; i <= mainArg.Range.End; i++ {
				iterationRes, err := solveNode(secondArg, []_VariableContext{
					{
						Name:  mainArg.Variable.Name,
						Value: float64(i),
					},
				})
				if err != nil {
					return 0, err
				}

				sum += iterationRes
			}

			return sum, nil
		}

	case lexemes.FUNCTION_PROD:
		{
			var prod float64 = 1
			for i := mainArg.Range.Start; i <= mainArg.Range.End; i++ {
				iterationRes, err := solveNode(secondArg, []_VariableContext{
					{
						Name:  mainArg.Variable.Name,
						Value: float64(i),
					},
				})
				if err != nil {
					return 0, err
				}

				prod *= iterationRes
			}

			return prod, nil
		}

	default:
		return 0, fmt.Errorf("undefined function \"%s\"", funcName)
	}
}

func solveNodeFuncCall(node parser.NodeFuncCall) (float64, error) {
	if lexemes.IsIRangeFunction(node.Name) {
		return solveNodeIRangeFuncCall(node)
	}

	return solveNodeDefaultFuncCall(node)
}

func solveNodeBinary(node parser.NodeBinary, varCtx []_VariableContext) (float64, error) {
	left, err := solveNode(node.Left, varCtx)
	if err != nil {
		return 0, err
	}

	right, err := solveNode(node.Right, varCtx)
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

func solveNode(node parser.Node, varCtx []_VariableContext) (float64, error) {

	switch typedNode := node.(type) {
	case parser.NodeNumber:
		return typedNode.Value, nil

	case parser.NodeVariable:
		{
			for _, ctx := range varCtx {
				if ctx.Name == typedNode.Name {
					return ctx.Value, nil
				}
			}

			return 0, fmt.Errorf("cannot find variable contenx")
		}

	case parser.NodeConstant:
		return solveNodeConstant(typedNode)

	case parser.NodeFuncCall:
		return solveNodeFuncCall(typedNode)

	case parser.NodeBinary:
		return solveNodeBinary(typedNode, varCtx)

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
	return solveNode(ast.Value, []_VariableContext{})
}

func EvaluateString(s string) (float64, error) {
	ast, err := CompileToAST(s)
	if err != nil {
		return 0, err
	}

	return EvaluateAST(ast)
}
