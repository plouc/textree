package textree

import (
	"fmt"
	"strings"
)

// RenderOptions is used to customize rendering
type RenderOptions struct {
	// symbols
	HorizontalLink string
	VerticalLink   string
	ChildLink      string
	LastChildLink  string
	ChildrenLink   string

	// dimensions
	LeafLinkLen         int
	ChildrenPaddingPre  int
	ChildrenPaddingPost int
}

// NewRenderOptions generates default rendering options
func NewRenderOptions() *RenderOptions {
	return &RenderOptions{
		// symbols
		HorizontalLink: "─",
		VerticalLink:   "│",
		ChildLink:      "├",
		LastChildLink:  "└",
		ChildrenLink:   "┬",

		// dimensions
		LeafLinkLen:         2,
		ChildrenPaddingPre:  1,
		ChildrenPaddingPost: 1,
	}
}

// Render renders a pretty tree structure
func (n *Node) Render(o *RenderOptions) string {
	line := ""

	reversedAncestors := n.ReversedAncestors()
	for _, ancestor := range reversedAncestors {
		if ancestor.isLast {
			line += "   "
		} else {
			line += "│  "
		}
	}

	if n.isLast {
		line += o.LastChildLink
	} else {
		line += o.ChildLink
	}
	if n.HasChild() {
		line += strings.Repeat(o.HorizontalLink, 2)
		line += o.ChildrenLink
	} else {
		line += strings.Repeat(o.HorizontalLink, 3)
	}
	line += fmt.Sprintf(" %s", n.Label)

	lines := []string{line}

	if n.HasChild() && o.ChildrenPaddingPre > 0 {
		childrenPrePadding := ""
		for _, ancestor := range n.Children[0].ReversedAncestors() {
			if ancestor.isLast {
				childrenPrePadding += "   "
			} else {
				childrenPrePadding += "│  "
			}
		}
		childrenPrePadding += "│"

		for i := 0; i < o.ChildrenPaddingPre; i++ {
			lines = append(lines, childrenPrePadding)
		}
	}

	for _, c := range n.Children {
		lines = append(lines, c.Render(o))
	}

	if n.HasChild() && o.ChildrenPaddingPost > 0 {
		for i := 0; i < o.ChildrenPaddingPost; i++ {
			lines = append(lines, "")
		}
	}

	return strings.Join(lines, "\n")
}
