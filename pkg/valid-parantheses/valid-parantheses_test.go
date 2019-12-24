package valid_parantheses

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValid(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"simple test 1": {"()", true},
		"simple test 2": {"()[]{}", true},
		"simple test 3": {"(]", false},
		"simple test 4": {"([)]", false},
		"simple test 5": {"{[]}", true},
		"empty input":   {"", true},
		"invert":        {"}{", false},
		"single 1":      {"(", false},
		"single 2":      {")", false},
		"different":     {"{]", false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsValid(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
