package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/singleton"
)

func main() {
	founder := singleton.GetFounder()
	var founderAge uint = 10
	founderName := "magic name"
	founder.SetAge(founderAge)
	founder.SetName(founderName)
	fmt.Println(founder)
}
