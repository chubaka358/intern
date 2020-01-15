package linked_list

// listNode is definition for singly-linked list.
type listNode struct {
	val  int
	next ListNoder
}

// SetNextNode sets next node for listNode
func (l *listNode) SetNextNode(list ListNoder) ListNoder {
	l.next = list
	return l.next
}

func (l *listNode) GetNext() ListNoder {
	return l.next
}

// NewListNode returns new listNode
func NewListNode(val int) ListNoder {
	return &listNode{
		val:  val,
		next: nil,
	}
}
