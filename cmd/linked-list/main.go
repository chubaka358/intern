package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/linked-list"
)

func main() {
	fmt.Println(linked_list.HasCycle(&linked_list.ListNode{
		Val: 1,
		Next: &linked_list.ListNode{
			Val: 2,
			Next: &linked_list.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}))
}
