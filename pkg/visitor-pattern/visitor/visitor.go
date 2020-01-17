package visitor

import "github.com/jayhrat/intern/pkg/visitor-pattern/transportation"

// Visitor provides a visitor-pattern interface
type Visitor interface {
	Visit(p transportation.Transport) string
}

// visitor implements Visitor interface
type visitor struct {
}

// Visit implements visit to airport
func (m *visitor) Visit(p transportation.Transport) string {
	return p.Use()
}

// NewMan creates and returns new visitor
func NewMan() Visitor {
	return &visitor{}
}
