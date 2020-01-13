# Mediator
Mediator pattern example
## Usage
```go
package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/mediator"
)

func main() {
	button := mediator.NewButton()
	textField := mediator.NewTextField()
	checkbox := mediator.NewCheckbox()
	mediator := mediator.NewActionMediator()
	mediator.Connect(button, textField, checkbox)
	fmt.Println(textField.SendData())
	fmt.Println(checkbox.SendData())
	fmt.Println(button.SendData())
}
```
