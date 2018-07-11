package textree

import (
	"bytes"
	"github.com/plouc/gosnap"
	"testing"
)

func getSampleTree() *Node {
	root := NewNode("ROOT")

	node1 := NewNode("1")
	root.Append(node1)

	node11 := NewNode("1.1")
	node1.Append(node11)
	node11.Append(NewNode("1.1.1"))
	node11.Append(NewNode("1.1.2"))

	node12 := NewNode("1.2")
	node1.Append(node12)
	node12.Append(NewNode("1.2.1"))

	node13 := NewNode("1.3")
	node1.Append(node13)
	node13.Append(NewNode("1.3.1"))
	node132 := NewNode("1.3.2")
	node13.Append(node132)
	node132.Append(NewNode("1.3.2.1"))
	node132.Append(NewNode("1.3.2.2"))
	node13.Append(NewNode("1.3.3"))

	node1.Append(NewNode("1.4"))

	node1.Append(NewNode("1.5"))

	return root
}

var testCases = []struct {
	name             string
	tree             *Node
	getRenderOptions func() *RenderOptions
}{
	{
		"basic",
		getSampleTree(),
		func() *RenderOptions {
			o := NewRenderOptions()

			return o
		},
	},
	{
		"dotted",
		getSampleTree(),
		func() *RenderOptions {
			o := NewRenderOptions()
			o.Dotted()

			return o
		},
	},
	{
		"rounded",
		getSampleTree(),
		func() *RenderOptions {
			o := NewRenderOptions()
			o.Rounded()

			return o
		},
	},
}

func TestRender(t *testing.T) {
	ctx := gosnap.NewContext(t, "snapshots")

	for _, testCase := range testCases {
		s := ctx.NewSnapshot(testCase.name)

		buf := new(bytes.Buffer)
		testCase.tree.Render(buf, testCase.getRenderOptions())
		s.AssertString(buf.String())
	}
}
