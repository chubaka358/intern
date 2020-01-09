package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/singleton"
)

func main() {
	founder1 := singleton.NewFounder()
	founder2 := singleton.NewFounder()
	fmt.Println(founder1 == founder2)
}
