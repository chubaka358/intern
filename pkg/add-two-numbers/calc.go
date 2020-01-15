package add_two_numbers

// calc implements Calcer
type calc struct {
}

// AddTwoNumbers returns the result of adding two numbers
func (c *calc) AddTwoNumbers(l1 ListNoder, l2 ListNoder) ListNoder {
	head := NewListNode(0)
	p := l1
	q := l2
	curr := head
	carry := 0
	for p != nil || q != nil {
		sum := 0
		if p != nil {
			sum += p.GetValue()
		}
		if q != nil {
			sum += q.GetValue()
		}
		sum += carry
		carry = sum / 10
		curr.SetNextNode(NewListNode(sum % 10))
		curr = curr.GetNext()
		if p != nil {
			p = p.GetNext()
		}
		if q != nil {
			q = q.GetNext()
		}
	}
	if carry > 0 {
		curr.SetNextNode(NewListNode(1))
	}
	return head.GetNext()
}

// NewCalc return new calc
func NewCalc() Calcer {
	return &calc{}
}
