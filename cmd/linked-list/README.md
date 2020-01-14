# Linked List Cycle
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

## Usage
```go
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
```

### Link
https://leetcode.com/problems/linked-list-cycle/
