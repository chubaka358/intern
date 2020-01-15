package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/chain-of-responsibility"
)

func main() {
	handlers := chain_of_responsibility.NewFirstHandler(
		chain_of_responsibility.NewSecondHandler(
			chain_of_responsibility.NewThirdHandler(
				chain_of_responsibility.NewFourthHandler(
					chain_of_responsibility.NewTerminator()))))
	fmt.Println(handlers.SendRequest("data type 3"))
}
