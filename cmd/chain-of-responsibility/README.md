# Chain Of Responsibility

Chain Of Responsibility pattern example

## Usage
```go
package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/chain-of-responsibility"
)

func main() {
	handlers := chain_of_responsibility.GetHandler{&chain_of_responsibility.DeleteHandler{&chain_of_responsibility.PostHandler{&chain_of_responsibility.PutHandler{}}}}
	fmt.Println("%q", handlers.SendRequest("GT", "1"))
}
```
