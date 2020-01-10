package visitor

import "testing"

func TestVisitor(t *testing.T) {
	transport := NewTransport()
	transport.Add(NewAirport())
	transport.Add(NewSubway())
	transport.Add(NewTeleport())

	want := "using airport...using subway...using teleport..."
	got := transport.Accept(NewMan())
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
