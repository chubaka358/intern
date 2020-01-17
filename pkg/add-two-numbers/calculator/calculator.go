package calculator

// Calculator provides calculator interface
type Calculator interface {
	Sum(n1, n2 *Node) (head Node)
}

// calculator implements Calculator
type calculator struct {
}

// AddTwoNumbers returns the result of adding two numbers
func (c *calculator) Sum(n1, n2 *Node) (head Node) {
	curr := &head
	carry := 0
	for n1 != nil || n2 != nil {
		sum := 0
		if n1 != nil {
			sum += n1.Val
		}
		if n2 != nil {
			sum += n2.Val
		}
		sum += carry
		carry = sum / 10
		curr.Next = &Node{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	if carry > 0 {
		curr.Next = &Node{
			Val:  1,
			Next: nil,
		}
	}
	return *head.Next
}

// NewCalculator return new calculator
func NewCalculator() Calculator {
	return &calculator{}
}
