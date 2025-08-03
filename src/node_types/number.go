package node_types

type Number struct {
	BaseNode
	Value int
}

func (n *Number) GetChildren() []Node {
	return []Node{}
}
