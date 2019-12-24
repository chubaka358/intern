package singleton

import "testing"

func TestGetFounder(t *testing.T) {
	founder1 := GetFounder()
	founder2 := GetFounder()
	if founder1 != founder2 {
		t.Fatalf("objects are not equal\n")
	}
}
