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
	t1.Val += t2.Val
	t1.Left = m.MergeTrees(t1.Left, t2.Left)
	t1.Right = m.MergeTrees(t1.Right, t2.Right)
	return t1
}

// NewMerge returns merge
func NewMerge() Merger {
	return &merge{}
}
