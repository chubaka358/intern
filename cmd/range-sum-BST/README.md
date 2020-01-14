# Range Sum of BST
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.

## Usage
```
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
```
```go
package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/range_sum_BST"
)

func main() {
	tree := range_sum_BST.NewTreeNode(10,
		range_sum_BST.NewTreeNode(5,
			range_sum_BST.NewTreeNode(3, nil, nil),
			range_sum_BST.NewTreeNode(7, nil, nil)),
		range_sum_BST.NewTreeNode(15, nil, range_sum_BST.NewTreeNode(18, nil, nil)))
	fmt.Println(tree.RangeSumBST(7, 15))
}
```

## Link
https://leetcode.com/problems/range-sum-of-bst/
