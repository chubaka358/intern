package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/adapter"
)

func main() {
	adaptee := adapter.NewAdaptee()
	adapter := adapter.NewAdapter(adaptee)
	fmt.Println(adapter.SendJSON())
}
