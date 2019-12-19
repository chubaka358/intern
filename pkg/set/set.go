package set

type settable interface {
	Add(...string)
	Delete(string)
	Size() int
	Contains(string) bool
	Intersection(set) set
	Union(set) set
}

// set stores elements
type set map[string]bool

// NewSet creates and returns a new empty set
func NewSet() set {
	innerMap := make(map[string]bool)
	return innerMap
}

// Add adds s to the set
func (set set) Add(s ...string) {
	for _, element := range s {
		set[element] = true
	}
}

// Delete removes s from a set
func (set set) Delete(s string) {
	delete(set, s)
}

// Size returns the number of elements in the set
func (set set) Size() int {
	return len(set)
}

// Contains checks if the set contains s element or not
// Returns true if contains, else false
func (set set) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// Intersection returns a new set that is the intersection of set and anotherSet
func (set set) Intersection(anotherSet set) set {
	intersect := make(map[string]int)
	interSet := NewSet()
	for key := range set {
		intersect[key]++
	}
	for key := range anotherSet {
		intersect[key]++
	}
	for key, val := range intersect {
		if val == 2 {
			interSet.Add(key)
		}
	}
	return interSet
}

// Union returns a new set, which is the union of set and anotherSet
func (set set) Union(anotherSet set) set {
	unionSet := NewSet()
	for key := range set {
		if !unionSet.Contains(key) {
			unionSet.Add(key)
		}
	}
	for key := range anotherSet {
		if !unionSet.Contains(key) {
			unionSet.Add(key)
		}
	}
	return unionSet
}
