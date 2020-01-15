package add_two_numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	isEqual := func(want ListNoder, got ListNoder) bool {
		for want != nil && got != nil {
			if want.GetValue() != got.GetValue() {
				return false
			}
			want = want.GetNext()
			got = got.GetNext()
		}
		if want != got {
			return false
		}
		return true
	}
	t.Run("simple test", func(t *testing.T) {
		calc := NewCalc()
		list1 := NewListNode(2)
		list1.SetNextNode(NewListNode(4)).SetNextNode(NewListNode(3))
		list2 := NewListNode(5)
		list2.SetNextNode(NewListNode(6)).SetNextNode(NewListNode(4))
		want := NewListNode(7)
		want.SetNextNode(NewListNode(0)).SetNextNode(NewListNode(8))
		got := calc.AddTwoNumbers(list1, list2)
		if !isEqual(want, got) {
			t.Fatalf("want %v, got %v", want.GetValue(), got.GetValue())
		}
	})
	t.Run("hard test", func(t *testing.T) {
		calc := NewCalc()
		list1 := NewListNode(9)
		list1.SetNextNode(NewListNode(9))

		list2 := NewListNode(1)

		want := NewListNode(0)
		want.SetNextNode(NewListNode(0)).SetNextNode(NewListNode(1))
		got := calc.AddTwoNumbers(list1, list2)

		if !isEqual(want, got) {
			t.Fatalf("want %v, got %v", want.GetValue(), got.GetValue())
		}
	})
}
