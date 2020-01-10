package range_sum_BST

import "testing"

func TestRangeSumBST(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		tree := &treeNode{
			val: 10,
			left: &treeNode{
				val: 5,
				left: &treeNode{
					val:   3,
					left:  nil,
					right: nil,
				},
				right: &treeNode{
					val:   7,
					left:  nil,
					right: nil,
				},
			},
			right: &treeNode{
				val:  15,
				left: nil,
				right: &treeNode{
					val:   18,
					left:  nil,
					right: nil,
				},
			},
		}
		want := 32
		got := tree.RangeSumBST(7, 15)
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("example2", func(t *testing.T) {
		tree := &treeNode{
			val: 10,
			left: &treeNode{
				val: 5,
				left: &treeNode{
					val: 3,
					left: &treeNode{
						val:   1,
						left:  nil,
						right: nil,
					},
				},
				right: &treeNode{
					val: 7,
					left: &treeNode{
						val:   6,
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
			right: &treeNode{
				val: 15,
				left: &treeNode{
					val:   13,
					left:  nil,
					right: nil,
				},
				right: &treeNode{
					val:   18,
					left:  nil,
					right: nil,
				},
			},
		}
		want := 23
		got := tree.RangeSumBST(6, 10)
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
