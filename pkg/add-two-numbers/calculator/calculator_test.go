package calculator

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	isEqual := func(want *Node, got *Node) bool {
		for want != nil && got != nil {
			if want.Val != got.Val {
				return false
			}
			want = want.Next
			got = got.Next
		}
		if want != got {
			return false
		}
		return true
	}
	t.Run("simple test", func(t *testing.T) {
		calc := NewCalculator()
		//list1 := NewListNode(2)
		//list1.Set(NewListNode(4)).Set(NewListNode(3))
		list1 := &Node{
			Val: 2,
			Next: &Node{
				Val: 4,
				Next: &Node{
					Val:  3,
					Next: nil,
				},
			},
		}
		//list2 := NewListNode(5)
		//list2.Set(NewListNode(6)).Set(NewListNode(4))
		list2 := &Node{
			Val: 5,
			Next: &Node{
				Val: 6,
				Next: &Node{
					Val:  4,
					Next: nil,
				},
			},
		}
		//want := NewListNode(7)
		//want.Set(NewListNode(0)).Set(NewListNode(8))
		want := &Node{
			Val: 7,
			Next: &Node{
				Val: 0,
				Next: &Node{
					Val:  8,
					Next: nil,
				},
			},
		}
		got := calc.AddTwoNumbers(list1, list2)
		if !isEqual(want, got) {
			t.Fatalf("want %v, got %v", want.Val, got.Val)
		}
	})
	t.Run("hard test", func(t *testing.T) {
		calc := NewCalculator()
		//list1 := NewListNode(9)
		//list1.Set(NewListNode(9))
		list1 := &Node{
			Val: 9,
			Next: &Node{
				Val:  9,
				Next: nil,
			},
		}

		list2 := &Node{
			Val:  1,
			Next: nil,
		}

		//want := NewListNode(0)
		//want.Set(NewListNode(0)).Set(NewListNode(1))
		want := &Node{
			Val: 0,
			Next: &Node{
				Val: 0,
				Next: &Node{
					Val:  1,
					Next: nil,
				},
			},
		}
		got := calc.AddTwoNumbers(list1, list2)

		if !isEqual(want, got) {
			t.Fatalf("want %v, got %v", want.Val, got.Val)
		}
	})
}
