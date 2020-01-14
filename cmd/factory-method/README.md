# Payment factory
Creates and returns a new payment object
## Usage
```go
package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/factory-method"
)

func main() {
	cash := factory_method.NewCash()
	fmt.Println(cash)
}
```
