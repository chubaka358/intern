package range_sum_BST

import "testing"

func TestRangeSumBST(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		tree := &TreeNode{
			Val:   10,
			Left:  &TreeNode{
				Val:   5,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: &TreeNode{
					Val:   18,
					Left:  nil,
					Right: nil,
				},
			},
		}
		want := 32
		got := RangeSumBST(tree, 7, 15)
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("example2", func(t *testing.T) {
		tree := &TreeNode{
			Val:   10,
			Left:  &TreeNode{
				Val:   5,
				Left:  &TreeNode{
					Val:   3,
					Left:  &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   15,
				Left:  &TreeNode{
					Val:   13,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   18,
					Left:  nil,
					Right: nil,
				},
			},
		}
		want := 23
		got := RangeSumBST(tree, 6, 10)
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
