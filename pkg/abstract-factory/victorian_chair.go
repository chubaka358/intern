package abstract_factory

// victorianChair implements victorian chair
type victorianChair struct {
}

// SitDown sits on the chair
func (v *victorianChair) SitDown() string {
	return "Sitting on the victorian chair"
}
