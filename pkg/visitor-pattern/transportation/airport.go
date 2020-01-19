package transportation

import "github.com/jayhrat/intern/pkg/visitor-pattern/visitor"

// airport implements the Transport interface
type airport struct {
}

// Use implementation
func (a *airport) Use() string {
	return "using airport..."
}

// Accept implementation
func (a *airport) Accept(v visitor.Visitor) string {
	return v.Visit(a)
}

// NewAirport creates and returns new airport
func NewAirport() Transport {
	return &airport{}
}
