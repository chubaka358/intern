package abstract_factory

// victorianSofa implements victorian sofa
type victorianSofa struct {
}

// LieDown lying on the victorian sofa
func (v *victorianSofa) LieDown() string {
	return "Lying on the victorianSofa"
}
