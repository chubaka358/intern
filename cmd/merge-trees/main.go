package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/merge-trees"
)

func main() {
	tree1 := merge_trees.NewTreeNode()
	tree1.Val = 1
	tree2 := merge_trees.NewTreeNode()
	tree2.Val = 5

	fmt.Println(tree1.MergeTrees(tree1, tree2))
}
