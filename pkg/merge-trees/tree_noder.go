package merge_trees

// TreeNoder provides treeNode interface
type TreeNoder interface {
	GetValue() int
	GetLeft() *treeNode
	GetRight() *treeNode
	SetLeft(left *treeNode)
	SetRight(right *treeNode)
	SetValue(val int)
}
