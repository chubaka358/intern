package add_two_numbers

// ListNoder provides listNode interface
type ListNoder interface {
	AddTwoNumbers(l1 *listNode, l2 *listNode) *listNode
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

// AddTwoNumbers returns the result of adding two numbers
func (l *listNode) AddTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	head := &listNode{}
	p := l1
	q := l2
	curr := head
	carry := 0
	for p != nil || q != nil {
		sum := 0
		if p != nil {
			sum += p.val
		}
		if q != nil {
			sum += q.val
		}
		sum += carry
		carry = sum / 10
		curr.next = &listNode{
			val: sum % 10,
		}
		curr = curr.next
		if p != nil {
			p = p.next
		}
		if q != nil {
			q = q.next
		}
	}
	if carry > 0 {
		curr.next = &listNode{
			val: 1,
		}
	}
	return head.next
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
