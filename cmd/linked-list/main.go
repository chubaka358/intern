package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/linked-list"
)

func main() {
	list := linked_list.NewListNode(1)
	list.SetNextNode(linked_list.NewListNode(2)).SetNextNode(list)
	fmt.Println(list.HasCycle(list))
}
