package interpreter

import (
	"fmt"
	"gocalc/parser"
	"gocalc/scanner"
	"math"
)

// func solveNodeDefaultFuncCall(node parser.NodeFuncCall) (float64, error) {
// 	funcName := node.Name
// 	funcArgc := node.Argc()
// 	funcArgs := []float64{}

// 	for _, arg := range node.Arguments {
// 		f, err := solveNode(arg, []_VariableContext{})
// 		if err != nil {
// 			return 0, err
// 		}

// 		funcArgs = append(funcArgs, f)
// 	}

// 	switch funcName {
// 	case l.FUNCTION_SIN:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Sin(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_COS:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Cos(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_TAN:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Tan(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_ATAN:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Atan(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_ABS:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Abs(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_LOG:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Log10(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_LN:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Log(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_SQRT:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Sqrt(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_CBRT:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Cbrt(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_ROUND:
// 		{
// 			if funcArgc != 1 {
// 				return 0, fmt.Errorf("%s expected 1 argument, but got %d", funcName, funcArgc)
// 			}

// 			return math.Round(funcArgs[0]), nil
// 		}

// 	case l.FUNCTION_RAND:
// 		{
// 			if funcArgc != 0 {
// 				return 0, fmt.Errorf("%s expected 0 argument, but got %d", funcName, funcArgc)
// 			}

// 			return rand.Float64(), nil
// 		}

// 	default:
// 		return 0, fmt.Errorf("undefined function \"%s\"", funcName)
// 	}
// }

// func solveNodeIRangeFuncCall(node parser.NodeFuncCall) (float64, error) {
// 	funcName := node.Name
// 	funcArgc := node.Argc()

// 	if funcArgc != 2 {
// 		return 0, fmt.Errorf("%s expected 2 arguments, but got %d", funcName, funcArgc)
// 	}

// 	assArg, isAssign := node.Arguments[0].(parser.NodeAssign)
// 	if !isAssign {
// 		return 0, fmt.Errorf(
// 			"%s expected an assign expression as first argument, but got %s",
// 			funcName,
// 			node.Arguments[0].Type(),
// 		)
// 	}
// 	assArgRight, isRange := assArg.Right.(parser.NodeRange)
// 	if !isRange {
// 		return 0, fmt.Errorf("%s expected a range expression as right hand expression", funcName)
// 	}

// 	secondArg := node.Arguments[1]

// 	switch funcName {
// 	case l.FUNCTION_SUM:
// 		{
// 			var sum float64 = 0
// 			for i := assArgRight.Start; i <= assArgRight.End; i++ {
// 				iterationRes, err := solveNode(secondArg, []_VariableContext{
// 					{
// 						Name:  assArg.Left.Name,
// 						Value: float64(i),
// 					},
// 				})
// 				if err != nil {
// 					return sum, err
// 				}

// 				sum += iterationRes
// 			}

// 			return sum, nil
// 		}

// 	case l.FUNCTION_PROD:
// 		{
// 			var prod float64 = 1
// 			for i := assArgRight.Start; i <= assArgRight.End; i++ {
// 				iterationRes, err := solveNode(secondArg, []_VariableContext{
// 					{
// 						Name:  assArg.Left.Name,
// 						Value: float64(i),
// 					},
// 				})
// 				if err != nil {
// 					return prod, err
// 				}

// 				prod *= iterationRes
// 			}

// 			return prod, nil
// 		}

// 	default:
// 		return 0, fmt.Errorf("undefined function \"%s\"", funcName)
// 	}
// }

// func solveNodeBinary(node parser.NodeBinary, varCtx []_VariableContext) (float64, error) {
// 	left, err := solveNode(node.Left, varCtx)
// 	if err != nil {
// 		return 0, err
// 	}

// 	right, err := solveNode(node.Right, varCtx)
// 	if err != nil {
// 		return 0, err
// 	}

// 	switch node.Operator {
// 	case l.OPERATOR_ADD:
// 		return left + right, nil

// 	case l.OPERATOR_SUB:
// 		return left - right, nil

// 	case l.OPERATOR_MUL:
// 		return left * right, nil

// 	case l.OPERATOR_DIV:
// 		return left / right, nil

// 	case l.OPERATOR_POW:
// 		return math.Pow(left, right), nil

// 	default:
// 		return 0, fmt.Errorf("undefined operator \"%s\"", node.Operator)
// 	}
// }

type _VarCtx struct {
	Name  string
	Value float64
}

func solveNodeIdentifier(node parser.NodeIdentifier, varCtx []_VarCtx) (float64, error) {
	def, isConst := IsConstant(node.Name)
	if isConst {
		return def.Value, nil
	}

	for _, v := range varCtx {
		if node.Name == v.Name {
			return v.Value, nil
		}
	}

	return 0, fmt.Errorf("undefined varaible \"%s\"", node.Name)
}

// todo
func solveNodeFuncCall(node parser.NodeFuncCall, varCtx []_VarCtx) (float64, error) {
	argc := node.Argc()

	def0, isFunc0 := IsFunc0(node.Name)
	if isFunc0 {
		if argc != 0 {
			return 0, fmt.Errorf("%s expected 0 arguments, but got %d", node.Name, argc)
		}

		return def0.Fn(), nil
	}

	def1, isFunc1 := IsFunc1(node.Name)
	if isFunc1 {
		if argc != 1 {
			return 0, fmt.Errorf("%s expected 1 argument, but got %d", node.Name, argc)
		}

		arg, err := solveNode(node.Arguments[0], varCtx)
		if err != nil {
			return 0, err
		}

		return def1.Fn(arg), nil
	}

	return 0, fmt.Errorf("undefined function \"%s\"", node.Name)
}

// todo
func solveNodeBinary(node parser.NodeBinary, varCtx []_VarCtx) (float64, error) {
	left, err := solveNode(node.Left, varCtx)
	if err != nil {
		return 0, err
	}

	right, err := solveNode(node.Right, varCtx)
	if err != nil {
		return 0, err
	}

	switch node.Operator {
	case "+":
		return left + right, nil

	case "-":
		return left - right, nil

	case "*":
		return left * right, nil

	case "/":
		return left / right, nil

	case "^":
		return math.Pow(left, right), nil

	case "=":
		return 0, fmt.Errorf("assign operator is unsolvable")

	case "..":
		return 0, fmt.Errorf("range operator is unsolvable")

	default:
		return 0, fmt.Errorf("undefined operator \"%s\"", node.Operator)
	}
}

// todo
func solveNode(node parser.Node, varCtx []_VarCtx) (float64, error) {
	switch nt := node.(type) {
	case parser.NodeNumber:
		return nt.Value, nil

	case parser.NodeIdentifier:
		return solveNodeIdentifier(nt, varCtx)

	case parser.NodeFuncCall:
		return solveNodeFuncCall(nt, varCtx)

	case parser.NodeBinary:
		return solveNodeBinary(nt, varCtx)

	default:
		return 0, fmt.Errorf("undefined node \"%s\"", node.Type())
	}
}

//---------------------------------------------------------------------------//

func EvaluateAST(ast parser.NodeRoot) (float64, error) {
	return solveNode(ast.Expression, []_VarCtx{})
}

func EvaluateString(s string) (float64, error) {
	tl, err := scanner.Scan(s)
	if err != nil {
		return 0, err
	}

	ast, err := parser.Parse(tl)
	if err != nil {
		return 0, err
	}

	return EvaluateAST(ast)
}
