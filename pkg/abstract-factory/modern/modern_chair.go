package modern

// chair provides modernChair interface
type chair interface {
	SitDown() string
}

// modernChair implements modern chair
type modernChair struct {
}

// SitDown sits on the chair
func (m *modernChair) SitDown() string {
	return "Sitting on the modern chair"
}
