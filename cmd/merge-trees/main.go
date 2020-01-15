package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/merge-trees"
)

func main() {
	merger := merge_trees.NewMerge()
	tree1 := merge_trees.NewTreeNode()
	tree1.SetValue(1)
	tree2 := merge_trees.NewTreeNode()
	tree2.SetValue(5)

	fmt.Println(merger.MergeTrees(tree1, tree2))
}
