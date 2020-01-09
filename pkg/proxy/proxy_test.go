package proxy

import "testing"

func TestProxy(t *testing.T) {
	dao := NewDAOProxy()
	want := "Doing inner Get logic\n"
	got := dao.Get(0)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
