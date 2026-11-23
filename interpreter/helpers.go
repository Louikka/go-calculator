package interpreter

import (
	"fmt"
	"gocalc/parser"
)

func checkFuncCallArgc(node parser.NodeFuncCall, targetArgc int) error {
	argc := node.Argc()

	if argc != targetArgc {
		s := "s"
		if targetArgc == 1 {
			s = ""
		}

		return fmt.Errorf("%s expected %d argument%s, but got %d", node.Name, targetArgc, s, argc)
	}

	return nil
}

// Helper for common math functions with one float64 argument and one float46
// return value (sin, cos, abs, etc.).
func commonfunc1(node parser.NodeFuncCall, varCtx []_VarCtx, fn func(float64) float64) (float64, error) {
	err := checkFuncCallArgc(node, 1)
	if err != nil {
		return 0, err
	}

	arg, err := solveNode(node.Arguments[0], varCtx)
	if err != nil {
		return 0, fmt.Errorf("%s cannot solve its argument: %s", node.Name, err)
	}

	return fn(arg), nil
}

// Helper for common math functions with two float64 arguments and one float46
// return value (nth root, log, etc.).
func commonfunc2(node parser.NodeFuncCall, varCtx []_VarCtx, fn func(float64, float64) float64) (float64, error) {
	err := checkFuncCallArgc(node, 2)
	if err != nil {
		return 0, err
	}

	arg1, err := solveNode(node.Arguments[0], varCtx)
	if err != nil {
		return 0, fmt.Errorf("%s cannot solve its 1st argument: %s", node.Name, err)
	}

	arg2, err := solveNode(node.Arguments[1], varCtx)
	if err != nil {
		return 0, fmt.Errorf("%s cannot solve its 2nd argument: %s", node.Name, err)
	}

	return fn(arg1, arg2), nil
}

// Helper for iterated binary operation function (https://w.wiki/ubPh). fn is
// a cumulative function that takes total value and result of the current
// iteration.
func ibofunc(node parser.NodeFuncCall, varCtx []_VarCtx, fn func(float64, float64) float64) (float64, error) {
	err := checkFuncCallArgc(node, 2)
	if err != nil {
		return 0, err
	}

	ass, err := parser.ConvertToAssign(node.Arguments[0])
	if err != nil {
		return 0, fmt.Errorf("%s cannot solve its 1st argument: %s", node.Name, err)
	}

	assRange, err := parser.ConvertToRange(ass.Expr)
	if err != nil {
		return 0, fmt.Errorf("%s cannot solve 1st argument assignment expression: %s", node.Name, err)
	}

	var cum float64 = 0

	for i := assRange.Start; i <= assRange.End; i++ {
		// iteration computation result
		r, err := solveNode(node.Arguments[1], []_VarCtx{
			{
				Name:  ass.Var.Name,
				Value: float64(i),
			},
		})
		if err != nil {
			return cum, err
		}

		cum = fn(cum, r)
	}

	return cum, nil
}
