package xml

// XMLSender provides xmlSender interface
type XMLSender interface {
	Send() string
}

// xmlSender implements XMLSender
type xmlSender struct {
}

// Send sends XML-formatted data
func (a *xmlSender) Send() string {
	return "<plant><name>Coffee</name><origin>Ethiopia</origin><origin>Brazil</origin></plant>"
}

// NewXMLSender returns new xmlSender
func NewXMLSender() XMLSender {
	return &xmlSender{}
}
