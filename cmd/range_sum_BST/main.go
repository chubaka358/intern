package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/range_sum_BST"
)

func main() {
	tree := range_sum_BST.NewTreeNode(10,
		range_sum_BST.NewTreeNode(5,
			range_sum_BST.NewTreeNode(3, nil, nil),
			range_sum_BST.NewTreeNode(7, nil, nil)),
		range_sum_BST.NewTreeNode(15, nil, range_sum_BST.NewTreeNode(18, nil, nil)))
	fmt.Println(tree.RangeSumBST(7, 15))
}
