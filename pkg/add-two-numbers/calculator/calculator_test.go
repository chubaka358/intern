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

		got := calc.Sum(list1, list2)
		if !isEqual(want, &got) {
			t.Fatalf("want %v, got %v", want.Val, got.Val)
		}
	})
	t.Run("hard test", func(t *testing.T) {
		calc := NewCalculator()
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
		got := calc.Sum(list1, list2)

		if !isEqual(want, &got) {
			t.Fatalf("want %v, got %v", want.Val, got.Val)
		}
	})
}
