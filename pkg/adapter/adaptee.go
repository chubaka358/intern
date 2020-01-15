package adapter

// adaptee implements adaptee
type adaptee struct {
}

// SendXML returns XML-formatted data
func (a *adaptee) SendXML() string {
	return "<plant><name>Coffee</name><origin>Ethiopia</origin><origin>Brazil</origin></plant>"
}

// NewAdaptee returns new adaptee
func NewAdaptee() XMLer {
	return &adaptee{}
}
