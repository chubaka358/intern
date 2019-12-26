package merge_trees

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeTrees(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		tree1 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}

		tree2 := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}
		want := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:  5,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}
		got := MergeTrees(tree1, tree2)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("empty trees", func(t *testing.T) {
		got := MergeTrees(nil, nil)
		if got != nil {
			t.Fatalf("want %v, got %v", nil, got)
		}
	})

	t.Run("empty and non-emtpy tree", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}
		got1 := MergeTrees(tree, nil)
		got2 := MergeTrees(nil, tree)
		diff1 := cmp.Diff(tree, got1)
		diff2 := cmp.Diff(tree, got2)
		if diff1 != "" {
			t.Errorf(diff1)
		}
		if diff2 != "" {
			t.Errorf(diff2)
		}
	})
}
