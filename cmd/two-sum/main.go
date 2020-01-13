package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/two-sum"
)

func main() {
	twoSum := two_sum.NewTwoSum()
	fmt.Println(twoSum.TwoSum([]int{4, 5, 3, 1}, 8))
}
