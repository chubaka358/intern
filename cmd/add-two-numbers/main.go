package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/add-two-numbers"
)

func main() {
	list1 := add_two_numbers.NewListNode(1)
	list1.SetNextNode(add_two_numbers.NewListNode(2)).SetNextNode(add_two_numbers.NewListNode(3))

	list2 := add_two_numbers.NewListNode(6)
	list2.SetNextNode(add_two_numbers.NewListNode(0)).SetNextNode(add_two_numbers.NewListNode(1))

	result := list1.AddTwoNumbers(list1, list2)
	for ; result != nil; result = result.GetNext() {
		fmt.Print(result.GetValue())
		fmt.Print("->")
	}
	fmt.Print("<nil>")
}
