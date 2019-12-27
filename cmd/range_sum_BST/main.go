package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/range_sum_BST"
)

func main() {
	tree := &range_sum_BST.TreeNode{
		Val: 10,
		Left: &range_sum_BST.TreeNode{
			Val: 5,
			Left: &range_sum_BST.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &range_sum_BST.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &range_sum_BST.TreeNode{
			Val:  15,
			Left: nil,
			Right: &range_sum_BST.TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(range_sum_BST.RangeSumBST(tree, 7, 15))
}
