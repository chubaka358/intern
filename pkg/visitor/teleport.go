package visitor

// teleport implements the Place interface
type teleport struct {
}

// UseTeleport implementation
func (t *teleport) UseTeleport() string {
	return "using teleport..."
}

// Accept implementation
func (t *teleport) Accept(v Visitor) string {
	return v.VisitTeleport(t)
}

// NewTeleport creates and returns new teleport
func NewTeleport() *teleport {
	return &teleport{}
}
