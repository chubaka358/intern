package json

import (
	"strings"

	"github.com/basgys/goxml2json"
)

// xmlSender provides xmlSender interface
type xmlSender interface {
	Send() string
}

// JSONSender provides jsonSender interface
type JSONSender interface {
	Send() string
}

// jsonSender implements JSONSender
type jsonSender struct {
	xmlSender xmlSender
}

// Send returns JSON-formatted data
func (a *jsonSender) Send() string {
	json, _ := xml2json.Convert(strings.NewReader(a.xmlSender.Send()))
	return json.String()
}

// NewJSONSender returns new jsonSender for xmlSender
func NewJSONSender(adaptee xmlSender) JSONSender {
	return &jsonSender{
		xmlSender: adaptee,
	}
}
