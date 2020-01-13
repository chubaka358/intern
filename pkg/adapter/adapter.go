package adapter

import (
	"strings"

	"github.com/basgys/goxml2json"
)

// adapter implements adapter
type adapter struct {
	*adaptee
}

// SendJSON returns JSON-formatted data
func (a *adapter) SendJSON() string {
	json, _ := xml2json.Convert(strings.NewReader(a.adaptee.SendXML()))
	return json.String()
}

// NewAdapter returns new adapter for adaptee
func NewAdapter(adaptee *adaptee) *adapter {
	return &adapter{adaptee}
}
