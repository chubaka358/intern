package merge_trees

// treeNode is definition for a binary tree node.
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// GetValue returns treeNode val
func (t *treeNode) GetValue() int {
	return t.val
}

// GetLeft returns treeNode left
func (t *treeNode) GetLeft() *treeNode {
	return t.left
}

// GetRight returns treeNode right
func (t *treeNode) GetRight() *treeNode {
	return t.right
}

// SetLeft sets treeNode left
func (t *treeNode) SetLeft(left *treeNode) {
	t.left = left
}

// SetRight sets treeNode right
func (t *treeNode) SetRight(right *treeNode) {
	t.right = right
}

// SetValue sets treeNode val
func (t *treeNode) SetValue(val int) {
	t.val = val
}

// NewTreeNode returns new treeNode
func NewTreeNode() *treeNode {
	return &treeNode{}
}
