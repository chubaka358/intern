package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/adapter/json"
	"github.com/jayhrat/intern/pkg/adapter/xml"
)

func main() {
	xmlSender := xml.NewXMLSender()
	jsonSender := json.NewJSONSender(xmlSender)
	fmt.Println(jsonSender.Send())
}
