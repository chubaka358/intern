package reverse_linked_list

type ListNoder interface {
	ReverseList() *listNode
	SetNextNode(val int) *listNode
	GetValue() int
	GetNext() *listNode
}

// ListNode is a definition for singly-linked list.
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

func (l *listNode) GetValue() int {
	return l.val
}

func (l *listNode) GetNext() *listNode {
	return l.next
}

func (l *listNode) SetNextNode(val int) *listNode {
	l.next = NewListNode(val)
	return l.next
}

func NewListNode(val int) *listNode {
	return &listNode{val: val}
}
