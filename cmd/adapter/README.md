# Adapter

Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

Allows you to convert xml to json

## Usage
**Code example**
```go
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
```
