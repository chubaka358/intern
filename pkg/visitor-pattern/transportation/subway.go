package transportation

import "github.com/jayhrat/intern/pkg/visitor-pattern/visitor"

// subway implements the Transport interface
type subway struct {
}

// Use implementation
func (s *subway) Use() string {
	return "using subway..."
}

// Accept implementation
func (s *subway) Accept(v visitor.Visitor) string {
	return v.Visit(s)
}

// NewSubway creates and returns new subway
func NewSubway() Transport {
	return &subway{}
}
