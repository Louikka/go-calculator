package parser

type Node interface {
	Type() string
}

// invalid

type NodeInvalid struct {
	//
}

func (n NodeInvalid) Type() string {
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

func NewNodeNumber(v float64) NodeNumber {
	return NodeNumber{
		Value: v,
	}
}

func (n NodeNumber) Type() string {
	return "NUMBER"
}

// range

type NodeRange struct {
	Start int
	End   int
}

func NewNodeRange(start, end int) NodeRange {
	return NodeRange{
		Start: start,
		End:   end,
	}
}

func (n NodeRange) Type() string {
	return "RANGE"
}

// identifier

type NodeIdentifier struct {
	Name string
}

func NewNodeIdentifier(name string) NodeIdentifier {
	return NodeIdentifier{
		Name: name,
	}
}

func (n NodeIdentifier) Type() string {
	return "IDENTIFIER"
}

// function call

type NodeFuncCall struct {
	Name      string
	Arguments []Node
}

func NewNodeFuncCall(name string, args []Node) NodeFuncCall {
	return NodeFuncCall{
		Name:      name,
		Arguments: args,
	}
}

func (n NodeFuncCall) Type() string {
	return "FUNCTION_CALL"
}

// Returns number of arguments.
func (n NodeFuncCall) Argc() int {
	return len(n.Arguments)
}

// binary expression

type NodeBinary struct {
	Operator string
	Left     Node
	Right    Node
}

func NewNodeBinary(oper string, l, r Node) NodeBinary {
	return NodeBinary{
		Operator: oper,
		Left:     l,
		Right:    r,
	}
}

func (n NodeBinary) Type() string {
	return "BINARY"
}
