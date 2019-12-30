package reverse_linked_list

// ListNode is a definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverse a singly linked list.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}
