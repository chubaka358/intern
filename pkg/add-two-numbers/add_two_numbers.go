package add_two_numbers

// ListNoder provides listNode interface
type ListNoder interface {
	SetNextNode(list ListNoder) ListNoder
	GetValue() int
	GetNext() ListNoder
}

// listNode implements list node
type listNode struct {
	val  int
	next ListNoder
}

// GetValue returns listNode value
func (l *listNode) GetValue() int {
	return l.val
}

// GetNext returns listNode next
func (l *listNode) GetNext() ListNoder {
	return l.next
}

// SetNextNode sets next node for listNode
func (l *listNode) SetNextNode(list ListNoder) ListNoder {
	l.next = list
	return l.next
}

// NewListNode returns new listNode
func NewListNode(val int) ListNoder {
	return &listNode{
		val:  val,
		next: nil,
	}
}
