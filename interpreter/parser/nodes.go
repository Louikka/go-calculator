package parser

type Node interface {
	Type() string
}

// invalid

type InvalidNode struct {
	//
}

func (n InvalidNode) Type() string {
	return "INVALID"
}

// root

type NodeRoot struct {
	Value Node
}

func (n NodeRoot) Type() string {
	return "ROOT"
}

// number

type NodeNumber struct {
	Value float64
}

func (n NodeNumber) Type() string {
	return "NUMBER"
}

// range

type NodeRange struct {
	Start int
	End   int
}

func (n NodeRange) Type() string {
	return "RANGE"
}

// variable

type NodeVariable struct {
	Name string
}

func (n NodeVariable) Type() string {
	return "VARIABLE"
}

// constant

type NodeConstant struct {
	Name string
}

func (n NodeConstant) Type() string {
	return "CONSTANT"
}

// function call

type NodeFuncCall struct {
	Name      string
	Arguments []Node
}

func (n NodeFuncCall) Type() string {
	return "FUNCTION_CALL"
}

// Returns arguments count.
//
//	len(node.Arguments)
func (n NodeFuncCall) Argc() int {
	return len(n.Arguments)
}

// binary expression

type NodeBinary struct {
	Operator string
	Left     Node
	Right    Node
}

func (n NodeBinary) Type() string {
	return "BINARY"
}

// assign expression

type NodeAssign struct {
	Left  NodeVariable
	Right Node
}

func (n NodeAssign) Type() string {
	return "ASSIGN"
}
