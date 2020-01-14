package reverse_linked_list

import (
	"testing"
)

func TestReverseList(t *testing.T) {

	isListsEqual := func(list1 *listNode, list2 *listNode) bool {
		currentNode1 := list1
		currentNode2 := list2
		for currentNode1 != nil && currentNode2 != nil {
			if currentNode1.GetValue() != currentNode2.GetValue() {
				return false
			}
			currentNode1 = currentNode1.next
			currentNode2 = currentNode2.next
		}
		if currentNode1 != currentNode2 {
			return false
		}
		return true
	}

	list1 := NewListNode(1)
	list1.SetNextNode(2).SetNextNode(3).SetNextNode(4).SetNextNode(5)
	list2 := NewListNode(5)
	list2.SetNextNode(4).SetNextNode(3).SetNextNode(2).SetNextNode(1)

	tests := map[string]struct {
		input *listNode
		want  *listNode
	}{
		"example 1": {list1, list2},
		"single element": {
			NewListNode(1), NewListNode(1),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.ReverseList()
			if !isListsEqual(got, tc.want) {
				t.Fatalf("want %#v, got %#v", tc.want, got)
			}
		})
	}
}
