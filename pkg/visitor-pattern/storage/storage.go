package storage

import (
	"github.com/jayhrat/intern/pkg/visitor-pattern/transportation"
)

type visitor interface {
	Visit(p transportation.Transport) string
}

type Storage interface {
	Add(p transportation.Transport)
	Accept(v visitor) string
}

// transportation implements a collection of places to visit
type storage struct {
	places []transportation.Transport
}

// Add appends Transport to the collection
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

// NewTransport creates and returns new transportation
func NewStorage() Storage {
	return &storage{}
}
