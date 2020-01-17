package visitor

import "github.com/jayhrat/intern/pkg/visitor-pattern/transportation"

// Visitor provides a visitor interface
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

// NewVisitor creates and returns new visitor-pattern
func NewVisitor() Visitor {
	return &visitor{}
}
