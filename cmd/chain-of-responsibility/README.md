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
	handlers := chain_of_responsibility.NewFirstHandler().
		SetNextHandler(chain_of_responsibility.NewSecondHandler().
			SetNextHandler(chain_of_responsibility.NewThirdHandler().
				SetNextHandler(chain_of_responsibility.NewFourthHandler())))
	fmt.Println(handlers.SendRequest("data type 3"))
}
```
