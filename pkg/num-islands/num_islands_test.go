package num_islands

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumIslands(t *testing.T) {
	tests := map[string]struct {
		input [][]byte
		want  int
	}{
		"example 1": {[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
			1},
		"example 2": {[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
			3},
		"zero islands": {[][]byte{
			{'0'},
		}, 0},
		"single island": {[][]byte{
			{'1'},
		}, 1},
		"hard example 1": {[][]byte{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		}, 1},
		"hard example 2": {[][]byte{
			{'1', '0', '1'},
			{'1', '1', '1'},
			{'1', '0', '1'},
		}, 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			islandsCounter := NewIslandsCounter()
			got := islandsCounter.NumIslands(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
