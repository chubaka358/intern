package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/add-two-numbers/calculator"
)

func main() {
	calc := calculator.NewCalculator()

	list1 := &calculator.Node{
		Val: 1,
		Next: &calculator.Node{
			Val: 2,
			Next: &calculator.Node{
				Val:  3,
				Next: nil,
			},
		},
	}

	list2 := &calculator.Node{
		Val: 6,
		Next: &calculator.Node{
			Val: 0,
			Next: &calculator.Node{
				Val:  1,
				Next: nil,
			},
		},
	}

	result := calc.Sum(list1, list2)
}
