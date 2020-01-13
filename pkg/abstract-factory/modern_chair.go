package abstract_factory

// modernChair implements modern chair
type modernChair struct {
}

// SitDown sits on the chair
func (m *modernChair) SitDown() string {
	return "Sitting on the modern chair"
}
