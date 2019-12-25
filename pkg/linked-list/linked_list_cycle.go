package linked_list

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle determine if list has a cycle in it
// Return true if has, else return false
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
