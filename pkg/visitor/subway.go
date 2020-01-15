package visitor

// subway implements the Place interface
type subway struct {
}

// UseSubway implementation
func (s *subway) UseSubway() string {
	return "using subway..."
}

// Accept implementation
func (s *subway) Accept(v Visitor) string {
	return v.VisitSubway(s)
}

// NewSubway creates and returns new subway
func NewSubway() Place {
	return &subway{}
}
