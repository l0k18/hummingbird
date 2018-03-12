package main

import (
	"fmt"

	"github.com/l0k1verloren/go-wbbst/pkg/tree32"
)

func main() {
	tree := tree32.NewTree()
	fmt.Println(tree)
	tree.AddRow()
	fmt.Println(tree)
	tree.AddRow()
	fmt.Println(tree)
	tree.AddRow()
	fmt.Println(tree)

}
