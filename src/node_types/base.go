package node_types

type Node interface {
	GetChildren() []Node
	GetStart() Location
	GetEnd() Location
}

type BaseNode struct {
	Children []Node
	Start    Location
	End      Location
}

func (n *BaseNode) GetChildren() []Node { return n.Children }
func (n *BaseNode) GetStart() Location  { return n.Start }
func (n *BaseNode) GetEnd() Location    { return n.End }

type Location struct {
	Line   int
	Column int
	Offset int
}
