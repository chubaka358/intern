# Visitor

Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

Allows visit airport, teleport and subway
## Usage
```go
package main

import (
	"fmt"
	
	"github.com/chubaka358/intern/pkg/visitor"
)

func main() {
	transport := visitor.NewTransport() // transport contains all places
	transport.Add(visitor.NewAirport())
	transport.Add(visitor.NewSubway())
	transport.Add(visitor.NewTeleport())

	fmt.Println(transport.Accept(visitor.NewMan()))
}
```
