package range_sum_BST

// TreeNoder provides treeNode interface
type TreeNoder interface {
	RangeSumBST(L int, R int) int
}

// TreeNode is definition for a binary tree node.
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// RangeSumBST return the sum of values of all nodes with value between L and R (inclusive)
func (t *treeNode) RangeSumBST(L int, R int) int {
	rangeSum := 0
	t.dfs(L, R, &rangeSum)
	return rangeSum
}

// dfs traverse tree
func (t *treeNode) dfs(L int, R int, rangeSum *int) {
	if L <= t.val && t.val <= R {
		*rangeSum += t.val
	}
	if t.val > L && t.left != nil {
		t.left.dfs(L, R, rangeSum)
	}
	if t.val < R && t.right != nil {
		t.right.dfs(L, R, rangeSum)
	}
}

// NewTreeNode creates new treeNode and returns pointer to treeNode
func NewTreeNode(val int, left *treeNode, right *treeNode) *treeNode {
	return &treeNode{
		val:   val,
		left:  left,
		right: right,
	}
}
