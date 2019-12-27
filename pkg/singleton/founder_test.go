package singleton

import "testing"

func TestGetFounder(t *testing.T) {
	founder1 := NewFounder()
	founder2 := NewFounder()
	if founder1 != founder2 {
		t.Fatalf("objects are not equal\n")
	}
}
