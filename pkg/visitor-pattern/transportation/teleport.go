package transportation

import "github.com/jayhrat/intern/pkg/visitor-pattern/visitor"

// teleport implements the Transport interface
type teleport struct {
}

// Use implementation
func (t *teleport) Use() string {
	return "using teleport..."
}

// Accept implementation
func (t *teleport) Accept(v visitor.Visitor) string {
	return v.Visit(t)
}

// NewTeleport creates and returns new teleport
func NewTeleport() Transport {
	return &teleport{}
}
