package chain_of_responsibility

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHandler_SendRequest(t *testing.T) {
	handlers := NewFirstHandler(NewSecondHandler(NewThirdHandler(NewFourthHandler(NewTerminator()))))
	t.Run("data type 1", func(t *testing.T) {
		want := "Using data type 1 handler"
		got := handlers.SendRequest("data type 1")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("data type 2", func(t *testing.T) {
		want := "Using data type 2 handler"
		got := handlers.SendRequest("data type 2")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("data type 3", func(t *testing.T) {
		want := "Using data type 3 handler"
		got := handlers.SendRequest("data type 3")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("data type 4", func(t *testing.T) {
		want := "Using data type 4 handler"
		got := handlers.SendRequest("data type 4")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("unknown data type", func(t *testing.T) {
		want := ""
		got := handlers.SendRequest("WTF")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})
}
