package two_sum

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		inputNums   []int
		inputTarget int
		want        []int
	}{
		"simple":           {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"two elements":     {[]int{1, 2}, 3, []int{0, 1}},
		"negative numbers": {[]int{-7, -5, -4, -1, -17}, -5, []int{2, 3}},
		"zero":             {[]int{0, -5, 5}, -5, []int{0, 1}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			twoSum := NewTwoSum()
			got := twoSum.TwoSum(tc.inputNums, tc.inputTarget)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
