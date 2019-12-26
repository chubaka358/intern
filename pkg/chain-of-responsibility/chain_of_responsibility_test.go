package chain_of_responsibility

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHandler_SendRequest(t *testing.T) {
	handlers := GetHandler{&DeleteHandler{&PostHandler{&PutHandler{}}}}
	t.Run("GET request", func(t *testing.T) {
		want := "Using GET method"
		got := handlers.SendRequest("GET", "")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("POST request", func(t *testing.T) {
		want := "Using POST method"
		got := handlers.SendRequest("POST", "")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("PUT request", func(t *testing.T) {
		want := "Using PUT method"
		got := handlers.SendRequest("PUT", "")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("DELETE request", func(t *testing.T) {
		want := "Using DELETE method"
		got := handlers.SendRequest("DELETE", "")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})

	t.Run("unknown request", func(t *testing.T) {
		want := ""
		got := handlers.SendRequest("WTF", "")
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Fatalf(diff)
		}
	})
}
