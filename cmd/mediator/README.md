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
	mediator.Connect(button, textField, checkbox)
	fmt.Println(textField.SendData()) // data was filled
	fmt.Println(checkbox.SendData()) // checkbox was checked
	fmt.Println(button.SendData()) // form has been submitted
}
```
