package abstract_factory

// modernSofa implements modern sofa
type modernSofa struct {
}

// LieDown lying on the modern sofa
func (s *modernSofa) LieDown() string {
	return "Lying on the modernSofa"
}
