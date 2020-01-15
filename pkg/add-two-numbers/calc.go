package add_two_numbers

// calc implements Calcer
type calc struct {
}

// AddTwoNumbers returns the result of adding two numbers
func (c *calc) AddTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
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

// NewCalc return new calc
func NewCalc() Calcer {
	return &calc{}
}
