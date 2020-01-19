package main

import (
	"fmt"
	"github.com/jayhrat/intern/pkg/abstract-factory/modern"
)

func main() {
	factory := modern.NewModernFurnitureFactory()
	sofa := factory.CreateSofa()
	fmt.Println(sofa.LieDown())
}
