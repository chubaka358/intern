package visitor_pattern

import (
	"testing"

	"github.com/jayhrat/intern/pkg/visitor-pattern/storage"
	"github.com/jayhrat/intern/pkg/visitor-pattern/transportation"
	"github.com/jayhrat/intern/pkg/visitor-pattern/visitor"
)

func TestVisitor(t *testing.T) {
	storage := storage.NewStorage()
	storage.Add(transportation.NewAirport())
	storage.Add(transportation.NewSubway())
	storage.Add(transportation.NewTeleport())

	want := "using airport...using subway...using teleport..."
	got := storage.Accept(visitor.NewMan())
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
