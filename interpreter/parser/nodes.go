package parser

type NodeType string

const (
	NODE_TYPE_ROOT          NodeType = "ROOT"
	NODE_TYPE_NUMBER        NodeType = "NUMBER"
	NODE_TYPE_VARIABLE      NodeType = "VARIABLE"
	NODE_TYPE_CONSTANT      NodeType = "CONSTANT"
	NODE_TYPE_FUNCTION_CALL NodeType = "FUNCTION_CALL"
	NODE_TYPE_EXPRESSION    NodeType = "EXPRESSION"
	NODE_TYPE_BINARY        NodeType = "BINARY"
)

type Node struct {
	Type  NodeType
	Value any
}

type NodeValueNumber struct {
	Value float64
}

type NodeValueVariable struct {
	Name string
}

type NodeValueConstant struct {
	Name string
}

type NodeValueFunction struct {
	Name     string
	Argument Node
}

type NodeValueBinary struct {
	Operator string
	Left     Node
	Right    Node
}
