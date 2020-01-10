package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/num-islands"
)

func main() {
	fmt.Println(num_islands.NewIslandsCount().NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}))
}
