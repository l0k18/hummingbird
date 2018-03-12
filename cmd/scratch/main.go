package main

import (
	"fmt"

	"github.com/l0k1verloren/go-wbbst/pkg/tree"
)

func main() {
	tree := tree.NewTree()
	fmt.Println(tree.Store)
}
