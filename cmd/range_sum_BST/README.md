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
```

## Link
https://leetcode.com/problems/range-sum-of-bst/
