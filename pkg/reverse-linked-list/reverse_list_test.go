package reverse_linked_list

import (
	"testing"
)

func TestReverseList(t *testing.T) {

	isListsEqual := func(list1 *ListNode, list2 *ListNode) bool {
		currentNode1 := list1
		currentNode2 := list2
		for currentNode1 != nil && currentNode2 != nil {
			if currentNode1.Val != currentNode2.Val {
				return false
			}
			currentNode1 = currentNode1.Next
			currentNode2 = currentNode2.Next
		}
		if currentNode1 != currentNode2 {
			return false
		}
		return true
	}

	tests := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"example 1": {&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}, &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		}},
		"single element": {
			&ListNode{
				Val:  1,
				Next: nil,
			}, &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ReverseList(tc.input)
			if !isListsEqual(got, tc.want) {
				t.Fatalf("want %#v, got %#v", tc.want, got)
			}
		})
	}
}
