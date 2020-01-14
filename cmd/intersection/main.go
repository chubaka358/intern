package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/intersection"
)

func main() {
	fmt.Println(intersection.NewIntersecter().Intersect([]int{3, 2, 1, 3}, []int{3, 1, 2}))
}
