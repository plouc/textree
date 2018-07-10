package textree

type Node struct {
	Label     string  `json:"label" yaml:"label"`
	Children  []*Node `json:"children,omitempty" yaml:"children,omitempty"`
	ancestors []*Node
	isLast    bool
}

func reverseNodes(nodes []*Node) []*Node {
	reversed := []*Node{}
	for i := len(nodes) - 1; i >= 0; i-- {
		reversed = append(reversed, nodes[i])
	}

	return reversed
}

func NewNode(label string) *Node {
	return &Node{
		Label:  label,
		isLast: true,
	}
}

func (n *Node) Depth() int {
	return len(n.ancestors)
}

func (n *Node) Parent() *Node {
	if len(n.ancestors) == 0 {
		return nil
	}

	return n.ancestors[0]
}

func (n *Node) Ancestors() []*Node {
	return n.ancestors
}

func (n *Node) ReversedAncestors() []*Node {
	return reverseNodes(n.ancestors)
}

func (n *Node) HasChild() bool {
	return len(n.Children) > 0
}

func (n *Node) IsRoot() bool {
	return len(n.ancestors) == 0
}

func (n *Node) IsLeaf() bool {
	return !n.HasChild()
}

func (n *Node) Append(c *Node) {
	if c.Parent() != nil {
		panic("Cannot append node already having a parent!")
	}

	for _, child := range n.Children {
		child.isLast = false
	}

	c.ancestors = append(n.ancestors, c.ancestors...)
	c.ancestors = append([]*Node{n}, c.ancestors...)
	c.isLast = true

	n.Children = append(n.Children, c)
}
