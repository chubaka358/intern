package transportation

// subway implements the Transport interface
type subway struct {
}

// UseSubway implementation
func (s *subway) Use() string {
	return "using subway..."
}

// Accept implementation
func (s *subway) Accept(v visitor) string {
	return v.Visit(s)
}

// NewSubway creates and returns new subway
func NewSubway() Transport {
	return &subway{}
}
