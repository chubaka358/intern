package range_sum_BST

// TreeNode is definition for a binary tree node.
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// dfs traverse tree
func dfs(node *treeNode, L int, R int, rangeSum *int) {
	if node != nil {
		if L <= node.val && node.val <= R {
			*rangeSum += node.val
		}
		if node.val > L {
			dfs(node.left, L, R, rangeSum)
		}
		if node.val < R {
			dfs(node.right, L, R, rangeSum)
		}
	}
}

// RangeSumBST return the sum of values of all nodes with value between L and R (inclusive)
func RangeSumBST(root *treeNode, L int, R int) int {
	rangeSum := 0
	dfs(root, L, R, &rangeSum)
	return rangeSum
}

// NewTreeNode creates new treeNode and returns pointer to treeNode
func NewTreeNode(val int, left *treeNode, right *treeNode) *treeNode {
	return &treeNode{
		val:   val,
		left:  left,
		right: right,
	}
}
