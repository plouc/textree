package textree

// Node is used to represent a tree/branch/leaf
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

// NewNode creates a new tree node
func NewNode(label string) *Node {
	return &Node{
		Label:  label,
		isLast: true,
	}
}

// Depth computes the node's depth in the tree
func (n *Node) Depth() int {
	return len(n.ancestors)
}

// Parent returns the immediate node's parent
func (n *Node) Parent() *Node {
	if len(n.ancestors) == 0 {
		return nil
	}

	return n.ancestors[0]
}

// Ancestors returns all node's ancestors,
// from the nearest to the farthest
func (n *Node) Ancestors() []*Node {
	return n.ancestors
}

// ReversedAncestors returns all node's ancestors,
// from the farthest to the nearest
func (n *Node) ReversedAncestors() []*Node {
	return reverseNodes(n.ancestors)
}

// HasChild checks if node has child
func (n *Node) HasChild() bool {
	return len(n.Children) > 0
}

// IsRoot checks if node is a root node
func (n *Node) IsRoot() bool {
	return len(n.ancestors) == 0
}

// IsLeaf checks if node is a leaf node
func (n *Node) IsLeaf() bool {
	return !n.HasChild()
}

// Append appends a new child to the node
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
