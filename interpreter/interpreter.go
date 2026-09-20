package interpreter

import (
	"fmt"
	"gocalc/parser"
	"gocalc/scanner"
	"math"
	"math/rand/v2"
)

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

	return 0, fmt.Errorf("undefined varaible %q", node.Name)
}

func solveNodeFuncCall(node parser.NodeFuncCall, varCtx []_VarCtx) (float64, error) {
	fnname := node.Name
	argc := node.Argc()

	// helper for common functions without arguments
	commonfunc0 := func(fn func() float64) (float64, error) {
		if argc != 0 {
			return 0, fmt.Errorf("%s expected 0 arguments, but got %d", fnname, argc)
		}

		return fn(), nil
	}

	// helper for common functions with 1 argument
	commonfunc1 := func(fn func(float64) float64) (float64, error) {
		if argc != 1 {
			return 0, fmt.Errorf("%s expected 1 argument, but got %d", fnname, argc)
		}

		arg, err := solveNode(node.Arguments[0], varCtx)
		if err != nil {
			return 0, err
		}

		return fn(arg), nil
	}

	// helper for SUM/PROD functions
	// fn is a cumulative function that takes total value and result of
	// the current iteration
	rfunc2 := func(fn func(float64, float64) float64) (float64, error) {
		if argc != 2 {
			return 0, fmt.Errorf("%s expected 2 arguments, but got %d", fnname, argc)
		}

		ass, err := parser.ConvertToAssign(node.Arguments[0])
		if err != nil {
			return 0, fmt.Errorf("%s: first argument error: %s", fnname, err)
		}

		assRange, err := parser.ConvertToRange(ass.Expr)
		if err != nil {
			return 0, fmt.Errorf("%s: first argument expression assignment error: %s", fnname, err)
		}

		expr := node.Arguments[1]

		var cum float64 = 0
		for i := assRange.Start; i <= assRange.End; i++ {
			iterResult, err := solveNode(expr, []_VarCtx{
				{
					Name:  ass.Var.Name,
					Value: float64(i),
				},
			})
			if err != nil {
				return cum, err
			}

			cum = fn(cum, iterResult)
		}

		return cum, nil
	}

	switch fnname {
	case "SIN":
		return commonfunc1(math.Sin)

	case "COS":
		return commonfunc1(math.Cos)

	case "TAN":
		return commonfunc1(math.Tan)

	case "ATAN":
		return commonfunc1(math.Atan)

	case "ABS":
		return commonfunc1(math.Abs)

	case "LOG":
		return commonfunc1(math.Log10)

	case "LN":
		return commonfunc1(math.Log)

	case "SQRT":
		return commonfunc1(math.Sqrt)

	case "CBRT":
		return commonfunc1(math.Cbrt)

	case "ROUND":
		return commonfunc1(math.Round)

	case "RAND":
		return commonfunc0(rand.Float64)

	case "SUM":
		return rfunc2(func(cum, i float64) float64 {
			return cum + i
		})

	case "PROD":
		return rfunc2(func(cum, i float64) float64 {
			if cum == 0 {
				cum = 1
			}
			return cum * i
		})

	default:
		return 0, fmt.Errorf("undefined function %q", node.Name)
	}
}

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
		return 0, fmt.Errorf("undefined operator %q", node.Operator)
	}
}

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
		return 0, fmt.Errorf("undefined node %q", node.Type())
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
