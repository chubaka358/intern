package merge_trees

// treeNode is definition for a binary tree node.
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// NewTreeNode returns new treeNode
func NewTreeNode() *treeNode {
	return &treeNode{}
}
