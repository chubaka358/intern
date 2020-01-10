package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/reverse-linked-list"
)

func main() {
	list := reverse_linked_list.NewListNode(1)
	list.SetNextNode(2).SetNextNode(3).SetNextNode(4)
	reversedList := list.ReverseList()
	for head := reversedList; head != nil; head = head.GetNext() {
		fmt.Printf("%v-->", head.GetValue())
	}
	fmt.Println("<nil>")
}
