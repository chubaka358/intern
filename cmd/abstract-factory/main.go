package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/abstract-factory"
)

func main() {
	factory := abstract_factory.NewModernFurnitureFactory()
	sofa := factory.CreateSofa()
	fmt.Println(sofa.LieDown())
}
