package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	list1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	list1.Next.Next.Next.Next = list1.Next

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	list2.Next.Next = list2

	list3 := &ListNode{
		Val:  1,
		Next: nil,
	}

	tests := map[string]struct {
		listNode *ListNode
		want     bool
	}{
		"example1": {list1, true},
		"example2": {list2, true},
		"example3": {list3, false},
		"nil head": {nil, false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := HasCycle(tc.listNode)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
