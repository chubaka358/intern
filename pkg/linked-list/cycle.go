package linked_list

type cycle struct {
}

// HasCycle determine if list has a cycle in it
// Return true if has, else return false
func (c *cycle) HasCycle(head *listNode) bool {
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

func NewCycle() Cycler {
	return &cycle{}
}
