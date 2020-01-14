# Set

Set data structure implementation

## Usage

```go
package main

import (
	"fmt"

	"github.com/jayhrat/intern/pkg/set"
)

func main() {
	mySet := set.NewSet()
	mySet.Add("15", "13", "11", "11", "11")
	fmt.Println(mySet)
	mySet.Delete("13")
	fmt.Println(mySet)
}
```
