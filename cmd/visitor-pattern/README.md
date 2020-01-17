# Visitor

Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

Allows visit airport, teleport and subway
## Usage
```go
package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/visitor-pattern/storage"
	"github.com/jayhrat/intern/pkg/visitor-pattern/transportation"
	"github.com/jayhrat/intern/pkg/visitor-pattern/visitor"
)

func main() {
	transport := storage.NewStorage()
	transport.Add(transportation.NewAirport())
	transport.Add(transportation.NewSubway())
	transport.Add(transportation.NewTeleport())

	fmt.Println(transport.Accept(visitor.NewMan()))
}
```
