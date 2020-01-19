package adapter

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jayhrat/intern/pkg/adapter/json"
	"github.com/jayhrat/intern/pkg/adapter/xml"
)

func TestAdapter(t *testing.T) {
	xmlSender := xml.NewXMLSender()
	jsonSender := json.NewJSONSender(xmlSender)
	want := "{\"plant\": {\"name\": \"Coffee\", \"origin\": [\"Ethiopia\", \"Brazil\"]}}\n"
	got := jsonSender.Send()
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Fatalf(diff)
	}
}
