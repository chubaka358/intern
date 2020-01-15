package merge_trees

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeTrees(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		tree1 := &treeNode{
			Val: 1,
			Left: &treeNode{
				Val: 3,
				Left: &treeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &treeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}

		tree2 := &treeNode{
			Val: 2,
			Left: &treeNode{
				Val:  1,
				Left: nil,
				Right: &treeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &treeNode{
				Val:  3,
				Left: nil,
				Right: &treeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}
		want := &treeNode{
			Val: 3,
			Left: &treeNode{
				Val: 4,
				Left: &treeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &treeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &treeNode{
				Val:  5,
				Left: nil,
				Right: &treeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}
		merger := NewMerge()
		got := merger.MergeTrees(tree1, tree2)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("empty trees", func(t *testing.T) {
		merger := NewMerge()
		got := merger.MergeTrees(nil, nil)
		if got != nil {
			t.Fatalf("want %v, got %v", nil, got)
		}
	})

	t.Run("empty and non-emtpy tree", func(t *testing.T) {
		tree := &treeNode{
			Val: 1,
			Left: &treeNode{
				Val: 3,
				Left: &treeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &treeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}
		merger := NewMerge()
		got1 := merger.MergeTrees(tree, nil)
		got2 := merger.MergeTrees(nil, tree)
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
