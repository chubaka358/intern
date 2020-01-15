package add_two_numbers

// ListNoder provides listNode interface
type ListNoder interface {
	SetNextNode(list *listNode) *listNode
	GetValue() int
	GetNext() *listNode
}

// listNode implements list node
type listNode struct {
	val  int
	next *listNode
}

// GetValue returns listNode value
func (l *listNode) GetValue() int {
	return l.val
}

// GetNext returns listNode next
func (l *listNode) GetNext() *listNode {
	return l.next
}

// SetNextNode sets next node for listNode
func (l *listNode) SetNextNode(list *listNode) *listNode {
	l.next = list
	return l.next
}

// NewListNode returns new listNode
func NewListNode(val int) *listNode {
	return &listNode{val, nil}
}
