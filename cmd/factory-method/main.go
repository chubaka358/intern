package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/factory-method"
)

func main() {
	cash := factory_method.NewCash()
	fmt.Println(cash)
}
