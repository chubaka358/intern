package linked_list

// cycle implements cycler
type cycle struct {
}

// HasCycle determine if list has a cycle in it
// Return true if has, else return false
func (c *cycle) HasCycle(head ListNoder) bool {
	if head == nil || head.GetNext() == nil {
		return false
	}
	slow := head
	fast := head.GetNext()
	for slow != fast {
		if fast == nil || fast.GetNext() == nil {
			return false
		}
		slow = slow.GetNext()
		fast = fast.GetNext().GetNext()
	}
	return true
}

// NewCycle returns new cycle
func NewCycle() Cycler {
	return &cycle{}
}
