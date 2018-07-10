package main

import (
	"fmt"
	"github.com/plouc/textree"
	"os"
	"strings"
)

func getSampleTree() *textree.Node {
	root := textree.NewNode("1")

	childA := textree.NewNode("1.1")
	root.Append(childA)
	childA.Append(textree.NewNode("1.1.1"))
	childA.Append(textree.NewNode("1.1.2"))

	childB := textree.NewNode("1.2")
	root.Append(childB)
	childB.Append(textree.NewNode("1.2.1"))

	childC := textree.NewNode("1.3")
	root.Append(childC)
	childC.Append(textree.NewNode("1.3.1"))
	childCChild := textree.NewNode("1.3.2")
	childC.Append(childCChild)
	childCChild.Append(textree.NewNode("1.3.2.1"))
	childCChild.Append(textree.NewNode("1.3.2.2"))
	childC.Append(textree.NewNode("1.3.3"))

	root.Append(textree.NewNode("1.4"))

	root.Append(textree.NewNode("1.5"))

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
	title("Dotted example", "using *RenderOptions.Dotted()")

	root := getSampleTree()

	o := textree.NewRenderOptions()
	o.Dotted()

	root.Render(os.Stdout, o)
}

func roundedExample() {
	title("Rounded example", "using *RenderOptions.Rounded()")

	root := getSampleTree()

	o := textree.NewRenderOptions()
	o.Rounded()

	root.Render(os.Stdout, o)
}

func compactExample() {
	title("Compact example", "using *RenderOptions.Compact()")

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
