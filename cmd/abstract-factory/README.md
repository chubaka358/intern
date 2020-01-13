#Abstract Factory
Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

##Usage
```go
package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/abstract-factory"
)

func main() {
	factory := abstract_factory.NewModernFurnitureFactory()
	sofa := factory.CreateSofa()
	fmt.Println(sofa.LieDown())
}
```
