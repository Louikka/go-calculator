package parser

import "fmt"

func ConvertToRange(n Node) (NodeRange, error) {
	r, isRange := n.(NodeRange)
	if isRange {
		return r, nil
	}

	bin, isBin := n.(NodeBinary)
	if !isBin {
		return NodeRange{}, ErrNotABinary
	}

	if bin.Operator != ".." {
		return NodeRange{}, fmt.Errorf("got \"%s\" operator instead of \"..\"", bin.Operator)
	}

	start, isNum := bin.Left.(NodeNumber)
	if !isNum || !start.IsInt() {
		return NodeRange{}, fmt.Errorf("expected an integer as start of the range")
	}

	end, isNum := bin.Right.(NodeNumber)
	if !isNum || !end.IsInt() {
		return NodeRange{}, fmt.Errorf("expected an integer as end of the range")
	}

	return NewNodeRange(int(start.Value), int(end.Value)), nil
}

func ConvertToAssign(n Node) (NodeAssign, error) {
	a, isAssign := n.(NodeAssign)
	if isAssign {
		return a, nil
	}

	bin, isBin := n.(NodeBinary)
	if !isBin {
		return NodeAssign{}, ErrNotABinary
	}

	if bin.Operator != "=" {
		return NodeAssign{}, fmt.Errorf("got \"%s\" operator instead of \"=\"", bin.Operator)
	}

	iden, isIden := bin.Left.(NodeIdentifier)
	if !isIden {
		return NodeAssign{}, fmt.Errorf("expected an identifier on the left of assign")
	}

	return NewNodeAssign(iden, bin.Right), nil
}
