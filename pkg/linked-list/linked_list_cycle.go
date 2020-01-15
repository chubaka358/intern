package linked_list

// listNode is definition for singly-linked list.
type listNode struct {
	val  int
	next *listNode
}

// HasCycle determine if list has a cycle in it
// Return true if has, else return false
func (l *listNode) HasCycle(head *listNode) bool {
	if head == nil || head.next == nil {
		return false
	}
	slow := head
	fast := head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
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
