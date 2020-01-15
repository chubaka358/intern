package merge_trees

// Merger provides merge interface
type Merger interface {
	MergeTrees(t1 *treeNode, t2 *treeNode) *treeNode
}
