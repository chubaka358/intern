package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/visitor"
)

func main() {
	transport := visitor.NewTransport()
	transport.Add(visitor.NewAirport())
	transport.Add(visitor.NewSubway())
	transport.Add(visitor.NewTeleport())

	fmt.Println(transport.Accept(visitor.NewMan()))
}
