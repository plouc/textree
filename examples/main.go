package main

import (
	"fmt"
	"github.com/plouc/textree"
	"os"
	"strings"
)

func getSampleTree() *textree.Node {
	root := textree.NewNode("ROOT")

	node1 := textree.NewNode("1")
	root.Append(node1)

	node11 := textree.NewNode("1.1")
	node1.Append(node11)
	node11.Append(textree.NewNode("1.1.1"))
	node11.Append(textree.NewNode("1.1.2"))

	node12 := textree.NewNode("1.2")
	node1.Append(node12)
	node12.Append(textree.NewNode("1.2.1"))

	node13 := textree.NewNode("1.3")
	node1.Append(node13)
	node13.Append(textree.NewNode("1.3.1"))
	node132 := textree.NewNode("1.3.2")
	node13.Append(node132)
	node132.Append(textree.NewNode("1.3.2.1"))
	node132.Append(textree.NewNode("1.3.2.2"))
	node13.Append(textree.NewNode("1.3.3"))

	node1.Append(textree.NewNode("1.4"))

	node1.Append(textree.NewNode("1.5"))

	return root
}

func title(t, d string) {
	fmt.Println(strings.Repeat("=-", 32))
	fmt.Println("")
	fmt.Printf("  %s\n", t)
	fmt.Printf("    %s\n", d)
	fmt.Println("")
	fmt.Println(strings.Repeat("=-", 32))
}

func basicExample() {
	title("Basic example", "using default options")

	root := getSampleTree()

	o := textree.NewRenderOptions()

	root.Render(os.Stdout, o)
}

func dottedExample() {
	title("Dotted example", "using RenderOptions.Dotted()")

	root := getSampleTree()

	o := textree.NewRenderOptions()
	o.Dotted()

	root.Render(os.Stdout, o)
}

func roundedExample() {
	title("Rounded example", "using RenderOptions.Rounded()")

	root := getSampleTree()

	o := textree.NewRenderOptions()
	o.Rounded()

	root.Render(os.Stdout, o)
}

func compactExample() {
	title("Compact example", "using RenderOptions.Compact()")

	root := getSampleTree()

	o := textree.NewRenderOptions()
	o.Compact()

	root.Render(os.Stdout, o)
}

func main() {
	basicExample()
	dottedExample()
	roundedExample()
	compactExample()
}
