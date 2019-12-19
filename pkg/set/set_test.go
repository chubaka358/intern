package set

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSet(t *testing.T) {
	var testSet = NewSet()

	if testSet == nil {
		t.Errorf("Expected set to be created")
	}
}

func TestSetAdd(t *testing.T) {
	t.Run("one argument", func(t *testing.T) {
		set := NewSet()
		value := "testValue"
		set.Add(value)
		if !set[value] {
			t.Errorf("Expected %v to be added in %v", value, set)
		}
	})

	t.Run("many arguments", func(t *testing.T) {
		set := NewSet()
		arguments := []string{"alala", "ololo", "trololo", "get", "cat"}
		set.Add(arguments...)
		for i := 0; i < len(arguments); i++ {
			if set.Contains(arguments[i]) == false {
				t.Errorf("Expected %v to be added in %v", arguments[i], set)
			}
		}
	})
}

func TestSetContains(t *testing.T) {
	set := NewSet()
	value := "containsValue"
	set.Add(value)
	want := true
	got := set.Contains(value)
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestSetDelete(t *testing.T) {
	set := NewSet()
	value := "valueToDelete"
	set.Add(value)
	set.Delete(value)
	want := false
	got := set.Contains(value)
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestSetSize(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		set := NewSet()
		want := 0
		got := set.Size()
		if want != got {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("nonempty set", func(t *testing.T) {
		set := NewSet()
		set.Add("Hello")
		set.Add("World")
		set.Add("Oh, shit")
		want := 3
		got := set.Size()
		if want != got {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

func TestSetIntersection(t *testing.T) {
	t.Run("two nonempty set", func(t *testing.T) {
		firstSet := NewSet()
		secondSet := NewSet()
		firstSet.Add("1", "2", "3", "4", "5")
		secondSet.Add("3", "6", "7", "1")
		want := NewSet()
		want.Add("1", "3")
		got := firstSet.Intersection(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("two empty set", func(t *testing.T) {
		firstSet := NewSet()
		secondSet := NewSet()
		want := NewSet()
		got := firstSet.Intersection(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("empty and nonempty set", func(t *testing.T) {
		firstSet := NewSet()
		secondSet := NewSet()
		secondSet.Add("3", "6", "7", "1")
		want := NewSet()
		got := firstSet.Intersection(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})
}

func TestSetUnion(t *testing.T) {
	t.Run("two empty set", func(t *testing.T) {
		firstSet := NewSet()
		secondSet := NewSet()
		want := NewSet()
		got := firstSet.Union(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("emty and nonempty", func(t *testing.T) {
		firstSet := NewSet()
		secondSet := NewSet()
		secondSet.Add("1", "2", "3", "4", "5")
		want := NewSet()
		want.Add("5", "4", "3", "2", "1")
		got := firstSet.Union(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("two nonempty", func(t *testing.T) {
		firstSet := NewSet()
		firstSet.Add("1", "7", "3", "15")
		secondSet := NewSet()
		secondSet.Add("1", "2", "3", "4", "5")
		want := NewSet()
		want.Add("5", "4", "3", "2", "1", "7", "15")
		got := firstSet.Union(secondSet)
		diff := cmp.Diff(want, got)
		if diff != "" {
			t.Errorf(diff)
		}
	})
}
