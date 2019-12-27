package range_sum_BST

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var rangeSum = 0

func RangeSumBST(root *TreeNode, L int, R int) int {
	rangeSum = 0
	dfs(root, L, R)
	return rangeSum
}

func dfs(node *TreeNode, L int, R int) {
	if node != nil {
		if L <= node.Val && node.Val <= R {
			rangeSum += node.Val
		}
		if node.Val > L {
			dfs(node.Left, L, R)
		}
		if node.Val < R {
			dfs(node.Right, L, R)
		}
	}
}
