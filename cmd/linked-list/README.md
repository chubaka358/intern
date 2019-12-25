# Linked List Cycle
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

## Usage
```go
package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/linked-list"
)

func main() {
	fmt.Println(linked_list.HasCycle(&linked_list.ListNode{
		Val:  1,
		Next: &linked_list.ListNode{
			Val:  2,
			Next: &linked_list.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})) // false
}
```

### Link
https://leetcode.com/problems/linked-list-cycle/
