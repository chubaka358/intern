package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/reverse-linked-list"
)

func main() {
	list := &reverse_linked_list.ListNode{
		Val: 5,
		Next: &reverse_linked_list.ListNode{
			Val: 3,
			Next: &reverse_linked_list.ListNode{
				Val:  17,
				Next: nil,
			},
		},
	}
	reversedList := reverse_linked_list.ReverseList(list)
	for head := reversedList; head != nil; head = head.Next {
		fmt.Printf("%v-->", head.Val)
	}
	fmt.Println("<nil>")
}
