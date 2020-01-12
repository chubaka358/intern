package visitor

// transport implements a collection of places to visit
type transport struct {
	places []Place
}

// Add appends Place to the collection
func (t *transport) Add(p Place) {
	t.places = append(t.places, p)
}

// Accept implements visit to all places in the transport
func (t *transport) Accept(v Visitor) string {
	var result string
	for _, place := range t.places {
		result += place.Accept(v)
	}
	return result
}

// NewTransport creates and returns new transport
func NewTransport() *transport {
	return &transport{}
}
