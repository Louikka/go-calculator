package interpreter

import (
	"fmt"
	l "gocalc/interpreter/lexemes"
	"gocalc/interpreter/parser"
	"gocalc/interpreter/scanner"
	"math"
	"math/rand/v2"
)

func solveNodeConstant(node parser.NodeConstant) (float64, error) {
	c, isConst := l.IsConstant(node.Name)

	if isConst {
		return c.Value, nil
	} else {
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
	case l.FUNCTION_SIN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Sin(funcArgs[0]), nil
		}

	case l.FUNCTION_COS:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Cos(funcArgs[0]), nil
		}

	case l.FUNCTION_TAN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Tan(funcArgs[0]), nil
		}

	case l.FUNCTION_ATAN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Atan(funcArgs[0]), nil
		}

	case l.FUNCTION_ABS:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Abs(funcArgs[0]), nil
		}

	case l.FUNCTION_LOG:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Log10(funcArgs[0]), nil
		}

	case l.FUNCTION_LN:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Log(funcArgs[0]), nil
		}

	case l.FUNCTION_SQRT:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Sqrt(funcArgs[0]), nil
		}

	case l.FUNCTION_CBRT:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Cbrt(funcArgs[0]), nil
		}

	case l.FUNCTION_ROUND:
		{
			if funcArgc != 1 {
				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
			}

			return math.Round(funcArgs[0]), nil
		}

	case l.FUNCTION_RAND:
		{
			if funcArgc != 0 {
				return 0, fmt.Errorf("%s expected 0 argument, but got %d", funcName, funcArgc)
			}

			return rand.Float64(), nil
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

	assArg, isAssign := node.Arguments[0].(parser.NodeAssign)
	if !isAssign {
		return 0, fmt.Errorf(
			"%s expected an assign expression as first argument, but got %s",
			funcName,
			node.Arguments[0].Type(),
		)
	}
	assArgRight, isRange := assArg.Right.(parser.NodeRange)
	if !isRange {
		return 0, fmt.Errorf("%s expected a range expression as right hand expression", funcName)
	}

	secondArg := node.Arguments[1]

	switch funcName {
	case l.FUNCTION_SUM:
		{
			var sum float64 = 0
			for i := assArgRight.Start; i <= assArgRight.End; i++ {
				iterationRes, err := solveNode(secondArg, []_VariableContext{
					{
						Name:  assArg.Left.Name,
						Value: float64(i),
					},
				})
				if err != nil {
					return sum, err
				}

				sum += iterationRes
			}

			return sum, nil
		}

	case l.FUNCTION_PROD:
		{
			var prod float64 = 1
			for i := assArgRight.Start; i <= assArgRight.End; i++ {
				iterationRes, err := solveNode(secondArg, []_VariableContext{
					{
						Name:  assArg.Left.Name,
						Value: float64(i),
					},
				})
				if err != nil {
					return prod, err
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
	if l.IsIRangeFunction(node.Name) {
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
	case l.OPERATOR_ADD:
		return left + right, nil

	case l.OPERATOR_SUB:
		return left - right, nil

	case l.OPERATOR_MUL:
		return left * right, nil

	case l.OPERATOR_DIV:
		return left / right, nil

	case l.OPERATOR_POW:
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
