package textree

import (
	"fmt"
	"io"
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

// Render renders a pretty tree structure to given io.Writer
func (n *Node) Render(w io.Writer, o *RenderOptions) {
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

	fmt.Fprintln(w, line)

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
			fmt.Fprintln(w, childrenPrePadding)
		}
	}

	for _, c := range n.Children {
		c.Render(w, o)
	}

	if n.HasChild() && o.ChildrenPaddingPost > 0 {
		for i := 0; i < o.ChildrenPaddingPost; i++ {
			fmt.Fprintln(w, "")
		}
	}
}
