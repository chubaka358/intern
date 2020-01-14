package merge_trees

// TreeNoder provides treeNode interface
type TreeNoder interface {
	MergeTrees(t1 *treeNode, t2 *treeNode) *treeNode
}

// treeNode is definition for a binary tree node.
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// MergeTrees merge two binary trees into a new binary tree
func (t *treeNode) MergeTrees(t1 *treeNode, t2 *treeNode) *treeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = t.MergeTrees(t1.Left, t2.Left)
	t1.Right = t.MergeTrees(t1.Right, t2.Right)
	return t1
}

// NewTreeNode returns new treeNode
func NewTreeNode() *treeNode {
	return &treeNode{}
}
