package node_types

type Expression struct {
	BaseNode
	Left  Node
	Right Node
}

func (e *Expression) GetChildren() []Node {
	return []Node{e.Left, e.Right}
}

type ExpressionStatement struct {
	BaseNode
	Expression Expression
}

func (e *ExpressionStatement) GetChildren() []Node {
	return []Node{&e.Expression}
}
