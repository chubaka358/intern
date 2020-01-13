package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	list1 := &listNode{
		Val: 3,
		Next: &listNode{
			Val: 2,
			Next: &listNode{
				Val: 0,
				Next: &listNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	list1.Next.Next.Next.Next = list1.Next

	list2 := &listNode{
		Val: 1,
		Next: &listNode{
			Val:  2,
			Next: nil,
		},
	}
	list2.Next.Next = list2

	list3 := &listNode{
		Val:  1,
		Next: nil,
	}

	tests := map[string]struct {
		listNode *listNode
		want     bool
	}{
		"example1": {list1, true},
		"example2": {list2, true},
		"example3": {list3, false},
		"nil head": {nil, false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			list := NewListNode(0)
			got := list.HasCycle(tc.listNode)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
