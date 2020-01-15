package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	list1 := &listNode{
		val: 3,
		next: &listNode{
			val: 2,
			next: &listNode{
				val: 0,
				next: &listNode{
					val:  -4,
					next: nil,
				},
			},
		},
	}
	list1.next.next.next.next = list1.next

	list2 := &listNode{
		val: 1,
		next: &listNode{
			val:  2,
			next: nil,
		},
	}
	list2.next.next = list2

	list3 := &listNode{
		val:  1,
		next: nil,
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
