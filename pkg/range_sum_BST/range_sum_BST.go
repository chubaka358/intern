package range_sum_BST

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// RangeSumBST return the sum of values of all nodes with value between L and R (inclusive)
func RangeSumBST(root *TreeNode, L int, R int) int {
	rangeSum := 0
	dfs(root, L, R, &rangeSum)
	return rangeSum
}

// dfs traverse tree
func dfs(node *TreeNode, L int, R int, rangeSum *int) {
	if node != nil {
		if L <= node.Val && node.Val <= R {
			*rangeSum += node.Val
		}
		if node.Val > L {
			dfs(node.Left, L, R, rangeSum)
		}
		if node.Val < R {
			dfs(node.Right, L, R, rangeSum)
		}
	}
}
