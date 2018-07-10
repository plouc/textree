# textree

[![GoDoc](https://godoc.org/github.com/plouc/textree?status.svg)](https://godoc.org/github.com/plouc/textree)
[![GitHub license](https://img.shields.io/github/license/plouc/textree.svg)](https://github.com/plouc/textree/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/plouc/textree)](https://goreportcard.com/report/github.com/plouc/textree)
[![GitHub issues](https://img.shields.io/github/issues/plouc/textree.svg)](https://github.com/plouc/textree/issues)

**textree** is a go package to easily pretty print nested trees in plain text.

## Install

```
go get github.com/plouc/textree
```

## Usage

```go
import (
	"os"
	"github.com/plouc/textree"
)
```

Construct a tree:

```go
root := textree.NewNode("1")

childA := textree.NewNode("1.1")
root.Append(childA)
childA.Append(textree.NewNode("1.1.1"))
childA.Append(textree.NewNode("1.1.2"))

childB := textree.NewNode("1.2")
root.Append(childB)
childB.Append(textree.NewNode("1.2.1"))
```

Then render the tree using the root element:

```go
o := textree.NewRenderOptions()
root.Render(os.Stdout, o)
```

Output:

```
└──┬ 1
   ├──┬ 1.1
   │  ├─── 1.1.1
   │  └─── 1.1.2
   └──┬ 1.2
      └─── 1.2.1
```

For complete usage of **textree**, see the full [package docs](https://godoc.org/github.com/plouc/textree).

## Examples

```
=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

  Basic example
    using default options

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
 └──┬ 1
    │
    ├──┬ 1.1
    │  │
    │  ├─── 1.1.1
    │  └─── 1.1.2
    │
    ├──┬ 1.2
    │  │
    │  └─── 1.2.1
    │
    ├──┬ 1.3
    │  │
    │  ├─── 1.3.1
    │  ├──┬ 1.3.2
    │  │  │
    │  │  ├─── 1.3.2.1
    │  │  └─── 1.3.2.2
    │  │
    │  └─── 1.3.3
    │
    ├─── 1.4
    └─── 1.5

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

  Dotted example
    using *RenderOptions.Dotted()

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
 ···· 1
    :
    :··· 1.1
    :  :
    :  :··· 1.1.1
    :  ···· 1.1.2
    :
    :··· 1.2
    :  :
    :  ···· 1.2.1
    :
    :··· 1.3
    :  :
    :  :··· 1.3.1
    :  :··· 1.3.2
    :  :  :
    :  :  :··· 1.3.2.1
    :  :  ···· 1.3.2.2
    :  :
    :  ···· 1.3.3
    :
    :··· 1.4
    ···· 1.5

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

  Rounded example
    using *RenderOptions.Rounded()

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
 ╰──╮ 1
    │
    ├──╮ 1.1
    │  │
    │  ├─── 1.1.1
    │  ╰─── 1.1.2
    │
    ├──╮ 1.2
    │  │
    │  ╰─── 1.2.1
    │
    ├──╮ 1.3
    │  │
    │  ├─── 1.3.1
    │  ├──╮ 1.3.2
    │  │  │
    │  │  ├─── 1.3.2.1
    │  │  ╰─── 1.3.2.2
    │  │
    │  ╰─── 1.3.3
    │
    ├─── 1.4
    ╰─── 1.5

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

  Compact example
    using *RenderOptions.Compact()

=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
└─┬ 1
  ├─┬ 1.1
  │ ├── 1.1.1
  │ └── 1.1.2
  ├─┬ 1.2
  │ └── 1.2.1
  ├─┬ 1.3
  │ ├── 1.3.1
  │ ├─┬ 1.3.2
  │ │ ├── 1.3.2.1
  │ │ └── 1.3.2.2
  │ └── 1.3.3
  ├── 1.4
  └── 1.5
```