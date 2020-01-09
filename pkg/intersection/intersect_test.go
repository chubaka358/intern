package intersection

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntersect(t *testing.T) {
	tests := map[string]struct {
		inputArg1 []int
		inputArg2 []int
		want      []int
	}{
		"Example 1": {[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		"Example 2": {[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		"empty":     {[]int{1, 2, 3, 3}, []int{}, []int{}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Intersect(tc.inputArg1, tc.inputArg2)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
