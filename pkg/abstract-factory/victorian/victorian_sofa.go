package victorian

// sofa provides victorianSofa interface
type sofa interface {
	LieDown() string
}

// victorianSofa implements victorian sofa
type victorianSofa struct {
}

// LieDown lying on the victorian sofa
func (v *victorianSofa) LieDown() string {
	return "Lying on the victorianSofa"
}
