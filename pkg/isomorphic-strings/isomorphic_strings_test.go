package isomorphic_strings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsomorphicStrings(t *testing.T) {
	tests := map[string]struct {
		inputS string
		inputT string
		want   bool
	}{
		"simpleTest1":                 {"egg", "add", true},
		"simpleTest2":                 {"foo", "bar", false},
		"simpleTest3":                 {"paper", "title", true},
		"simpleTest4":                 {"aa", "ab", false},
		"simpleTest5":                 {"ab", "aa", false},
		"simpleTest6":                 {"aab", "aaa", false},
		"simpleTest7":                 {"zz", "zz", true},
		"ChineseTest":                 {"汉字", "漢字", true},
		"emojiTest":                   {"😄😎😎", "🥏🥎🥎", true},
		"differentLength":             {"dog", "doge", false},
		"emptyStrings":                {"", "", true},
		"singleIdenticalCharacter":    {"a", "a", true},
		"singleNonidenticalCharacter": {"a", "b", true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			isomorph := NewIsomorph()
			got := isomorph.IsIsomorphic(tc.inputS, tc.inputT)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
