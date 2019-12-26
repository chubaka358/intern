package merge_trees

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MergeTrees merge two binary trees into a new binary tree
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	return t1
}
