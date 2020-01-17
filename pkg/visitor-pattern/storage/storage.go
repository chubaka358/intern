package storage

import "github.com/jayhrat/intern/pkg/visitor-pattern/transportation"

// visitor provides visitor interface
type visitor interface {
	Visit(p transportation.Transport) string
}

// Storage provides storage interface
type Storage interface {
	Add(p transportation.Transport)
	Accept(v visitor) string
}

// storage implements a collection of places to visit
type storage struct {
	places []transportation.Transport
}

// Add appends Transport to the storage
func (t *storage) Add(p transportation.Transport) {
	t.places = append(t.places, p)
}

// Accept implements visit to all places in the transportation
func (t *storage) Accept(v visitor) string {
	var result string
	for _, place := range t.places {
		result += place.Accept(v)
	}
	return result
}

// NewStorage creates and returns new storage
func NewStorage() Storage {
	return &storage{}
}
