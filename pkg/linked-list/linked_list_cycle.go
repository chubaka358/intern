package linked_list

// listNode is definition for singly-linked list.
type listNode struct {
	val  int
	next *listNode
}

// SetNextNode sets next node for listNode
func (l *listNode) SetNextNode(list *listNode) *listNode {
	l.next = list
	return l.next
}

// NewListNode returns new listNode
func NewListNode(val int) *listNode {
	return &listNode{
		val:  val,
		next: nil,
	}
}
