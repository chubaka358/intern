package merge_trees

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeTrees(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		tree1 := &treeNode{
			val: 1,
			left: &treeNode{
				val: 3,
				left: &treeNode{
					val:   5,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &treeNode{
				val:   2,
				left:  nil,
				right: nil,
			},
		}

		tree2 := &treeNode{
			val: 2,
			left: &treeNode{
				val:  1,
				left: nil,
				right: &treeNode{
					val:   4,
					left:  nil,
					right: nil,
				},
			},
			right: &treeNode{
				val:  3,
				left: nil,
				right: &treeNode{
					val:   7,
					left:  nil,
					right: nil,
				},
			},
		}
		want := &treeNode{
			val: 3,
			left: &treeNode{
				val: 4,
				left: &treeNode{
					val:   5,
					left:  nil,
					right: nil,
				},
				right: &treeNode{
					val:   4,
					left:  nil,
					right: nil,
				},
			},
			right: &treeNode{
				val:  5,
				left: nil,
				right: &treeNode{
					val:   7,
					left:  nil,
					right: nil,
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
			val: 1,
			left: &treeNode{
				val: 3,
				left: &treeNode{
					val:   5,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &treeNode{
				val:   2,
				left:  nil,
				right: nil,
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
