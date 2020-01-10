package visitor

type transport struct {
	places []Place
}

func (t *transport) Add(p Place) {
	t.places = append(t.places, p)
}

func (t *transport) Accept(v Visitor) string {
	var result string
	for _, place := range t.places {
		result += place.Accept(v)
	}
	return result
}

func NewTransport() *transport {
	return &transport{}
}
