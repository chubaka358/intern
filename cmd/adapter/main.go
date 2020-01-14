package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/adapter"
)

func main() {
	adaptee := adapter.NewAdaptee()
	adapter := adapter.NewAdapter(adaptee)
	fmt.Println(adapter.SendJSON())
}
