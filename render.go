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
	RootLink       string
	ChildLink      string
	LastChildLink  string
	ChildrenLink   string
	NodeSymbol     string

	// dimensions
	MarginTop            int
	MarginBottom         int
	MarginLeft           int
	HorizontalLinkLength int
	LabelPaddingLeft     int
	ChildrenMarginTop    int
	ChildrenMarginBottom int
}

// NewRenderOptions generates default rendering options
func NewRenderOptions() *RenderOptions {
	return &RenderOptions{
		// symbols
		HorizontalLink: "─",
		VerticalLink:   "│",
		RootLink:       "┌",
		ChildLink:      "├",
		LastChildLink:  "└",
		ChildrenLink:   "┬",
		NodeSymbol:     "",

		// dimensions
		MarginTop:            1,
		MarginBottom:         1,
		MarginLeft:           1,
		HorizontalLinkLength: 2,
		LabelPaddingLeft:     1,
		ChildrenMarginTop:    1,
		ChildrenMarginBottom: 1,
	}
}

// Rounded override symbols for a rounded rendering
func (o *RenderOptions) Rounded() {
	o.HorizontalLink = "─"
	o.VerticalLink = "│"
	o.ChildLink = "├"
	o.LastChildLink = "╰"
	o.ChildrenLink = "╮"
	o.NodeSymbol = ""
}

// Dotted override symbols for a dotted rendering
func (o *RenderOptions) Dotted() {
	o.HorizontalLink = "·"
	o.VerticalLink = ":"
	o.ChildrenLink = "·"
	o.ChildLink = ":"
	o.LastChildLink = "·"
}

// Compact override dimensions for a compact rendering
func (o *RenderOptions) Compact() {
	o.MarginTop = 0
	o.MarginBottom = 0
	o.MarginLeft = 0
	o.HorizontalLinkLength = 1
	o.LabelPaddingLeft = 1
	o.ChildrenMarginTop = 0
	o.ChildrenMarginBottom = 0
}

// Render renders a pretty tree structure to given io.Writer
func (n *Node) Render(w io.Writer, o *RenderOptions) {
	marginLeft := strings.Repeat(" ", o.MarginLeft)

	line := marginLeft

	if n.IsRoot() {
		fmt.Fprint(w, strings.Repeat("\n", o.MarginTop))
		line += o.RootLink
	} else {
		reversedAncestors := n.ReversedAncestors()
		for _, ancestor := range reversedAncestors {
			if ancestor.IsRoot() {
				continue
			}

			if ancestor.isLast {
				line += strings.Repeat(" ", o.HorizontalLinkLength+1)
			} else {
				line += o.VerticalLink
				line += strings.Repeat(" ", o.HorizontalLinkLength)
			}
		}

		if n.isLast {
			line += o.LastChildLink
		} else {
			line += o.ChildLink
		}
		if n.HasChild() {
			line += strings.Repeat(o.HorizontalLink, o.HorizontalLinkLength)
			line += o.ChildrenLink
		} else {
			line += strings.Repeat(o.HorizontalLink, o.HorizontalLinkLength+1)
		}
	}

	line += o.NodeSymbol
	line += strings.Repeat(" ", o.LabelPaddingLeft)
	line += n.Label

	fmt.Fprintln(w, line)

	if n.HasChild() && o.ChildrenMarginTop > 0 {
		childrenMarginTop := marginLeft
		for _, ancestor := range n.Children[0].ReversedAncestors() {
			if ancestor.IsRoot() {
				continue
			}
			if ancestor.isLast {
				childrenMarginTop += strings.Repeat(" ", o.HorizontalLinkLength+1)
			} else {
				childrenMarginTop += o.VerticalLink
				childrenMarginTop += strings.Repeat(" ", o.HorizontalLinkLength)
			}
		}
		childrenMarginTop += o.VerticalLink

		for i := 0; i < o.ChildrenMarginTop; i++ {
			fmt.Fprintln(w, childrenMarginTop)
		}
	}

	for _, c := range n.Children {
		c.Render(w, o)
	}

	if n.IsLeaf() && n.isLast && o.ChildrenMarginBottom > 0 {
		childrenMarginBottom := marginLeft
		for _, ancestor := range n.ReversedAncestors() {
			if ancestor.IsRoot() {
				continue
			}

			if ancestor.isLast {
				childrenMarginBottom += strings.Repeat(" ", o.HorizontalLinkLength+1)
			} else {
				childrenMarginBottom += o.VerticalLink
				childrenMarginBottom += strings.Repeat(" ", o.HorizontalLinkLength)
			}
		}

		for i := 0; i < o.ChildrenMarginBottom; i++ {
			fmt.Fprintln(w, childrenMarginBottom)
		}
	}

	if n.IsRoot() {
		fmt.Fprint(w, strings.Repeat("\n", o.MarginBottom))
	}
}
