package merge_trees

// merge implements Merger interface
type merge struct {
}

// MergeTrees merges t1 and t2
func (m *merge) MergeTrees(t1 *treeNode, t2 *treeNode) *treeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.val += t2.val
	t1.left = m.MergeTrees(t1.left, t2.left)
	t1.right = m.MergeTrees(t1.right, t2.right)
	return t1
}

// NewMerge returns merge
func NewMerge() Merger {
	return &merge{}
}
