package reverse_linked_list

// ListNoder provides listNode interface
type ListNoder interface {
	ReverseList() *listNode
	SetNextNode(val int) *listNode
	GetValue() int
	GetNext() *listNode
}

// listNode is a definition for singly-linked list.
type listNode struct {
	val  int
	next *listNode
}

// ReverseList reverse a singly linked list.
func (l *listNode) ReverseList() *listNode {
	var prev *listNode
	current := l
	for current != nil {
		nextTemp := current.next
		current.next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

// GetValue returns listNode val
func (l *listNode) GetValue() int {
	return l.val
}

// GetNext returns listNode next
func (l *listNode) GetNext() *listNode {
	return l.next
}

// SetNextNode sets next node for listNode and returns pointer to next node
func (l *listNode) SetNextNode(val int) *listNode {
	l.next = NewListNode(val)
	return l.next
}

// NewListNode returns new listNode
func NewListNode(val int) *listNode {
	return &listNode{val: val}
}
