package visitor

// airport implements the Place interface
type airport struct {
}

// UseAirport implementation
func (a *airport) UseAirport() string {
	return "using airport..."
}

// Accept implementation
func (a *airport) Accept(v Visitor) string {
	return v.VisitAirport(a)
}

// NewAirport creates and returns new airport
func NewAirport() Place {
	return &airport{}
}
