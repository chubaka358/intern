package adapter

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	want := "{\"plant\": {\"name\": \"Coffee\", \"origin\": [\"Ethiopia\", \"Brazil\"]}}\n"
	got := adapter.SendJSON()
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Fatalf(diff)
	}
}
