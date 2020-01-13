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
