package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	list1 := NewListNode(3)
	list1.SetNextNode(NewListNode(2)).
		SetNextNode(NewListNode(0)).
		SetNextNode(NewListNode(-4))
	list1.GetNext().GetNext().GetNext().SetNextNode(list1.GetNext())

	list2 := NewListNode(1)
	list2.SetNextNode(NewListNode(2))
	list2.GetNext().SetNextNode(list2)

	list3 := NewListNode(1)

	tests := map[string]struct {
		listNode ListNoder
		want     bool
	}{
		"example1": {list1, true},
		"example2": {list2, true},
		"example3": {list3, false},
		"nil head": {nil, false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cycle := NewCycle()
			got := cycle.HasCycle(tc.listNode)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
