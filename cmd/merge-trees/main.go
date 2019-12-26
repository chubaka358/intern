package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/merge-trees"
)

func main() {
	tree1 := &merge_trees.TreeNode{
		Val: 1,
		Left: &merge_trees.TreeNode{
			Val: 3,
			Left: &merge_trees.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &merge_trees.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	tree2 := &merge_trees.TreeNode{
		Val: 2,
		Left: &merge_trees.TreeNode{
			Val:  1,
			Left: nil,
			Right: &merge_trees.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &merge_trees.TreeNode{
			Val:  3,
			Left: nil,
			Right: &merge_trees.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	newTree := merge_trees.MergeTrees(tree1, tree2)
	fmt.Println(newTree)
}
