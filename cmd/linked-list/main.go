package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/linked-list"
)

func main() {
	list := linked_list.NewListNode(1, linked_list.NewListNode(2, linked_list.NewListNode(3, nil)))
	fmt.Println(linked_list.HasCycle(list))
}
