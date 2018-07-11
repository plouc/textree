package textree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	pathIDSeparator = "__"
)

// TreeFromDir generates a tree representing given directory structure
func TreeFromDir(dir string) (*Node, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	pathStat, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	if !pathStat.IsDir() {
		return nil, fmt.Errorf("path must be a valid directory, but %s isn't", absPath)
	}

	dirPathLen := len(absPath)

	// Manually creating root directory
	root := NewNode(dir)
	s := map[string]*Node{
		"": root,
	}

	walkErr := filepath.Walk(absPath, func(path string, f os.FileInfo, err error) error {
		// skipping root as it has already been defined
		if len(path) == dirPathLen {
			return nil
		}

		// removing base directory path + path separator
		rel := path[dirPathLen+1:]

		// splitting path
		parts := strings.Split(rel, string(os.PathSeparator))

		parent := root
		currentPath := ""
		for idx, x := range parts {
			if idx == len(parts)-1 {
				child := NewNode(x)
				parent.Append(child)
				s[currentPath+x+pathIDSeparator] = child
				continue
			}

			currentPath = currentPath + x + pathIDSeparator
			_, ok := s[currentPath]
			if !ok {
				s[currentPath] = NewNode(currentPath)
			}

			parent = s[currentPath]
		}

		return err
	})

	return root, walkErr
}
