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
func (s *set) Add(elements ...string) {
	for _, element := range elements {
		(*s)[element] = true
	}
}

// Delete removes s from a set
func (s *set) Delete(element string) {
	delete(*s, element)
}

// Size returns the number of elements in the set
func (s *set) Size() int {
	return len(*s)
}

// Contains checks if the set contains s element or not
// Returns true if contains, else false
func (s *set) Contains(element string) bool {
	_, ok := (*s)[element]
	return ok
}

// Intersection returns a new set that is the intersection of set and anotherSet
func (s *set) Intersection(anotherS *set) set {
	intersect := make(map[string]int)
	interSet := NewSet()
	for key := range *s {
		intersect[key]++
	}
	for key := range *anotherS {
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
func (s *set) Union(anotherS *set) set {
	unionSet := NewSet()
	for key := range *s {
		if !unionSet.Contains(key) {
			unionSet.Add(key)
		}
	}
	for key := range *anotherS {
		if !unionSet.Contains(key) {
			unionSet.Add(key)
		}
	}
	return unionSet
}
