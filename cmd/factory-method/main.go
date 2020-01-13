package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/factory-method"
)

func main() {
	cash := factory_method.NewCash()
	fmt.Println(cash)
}
