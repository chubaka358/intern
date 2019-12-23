# Payment factory
Creates and returns a new payment object
## Usage
```go
package main

import (
    "fmt"

    "github.com/chubaka358/intern/pkg/factory-method"
)

func main() {
    factory := factory_method.NewFactory() // create new factory
    
    creditCard, err := factory.CreatePayment("CreditCard")
    fmt.Println(err)
    creditCard.Replenish(1000) // add 1000 to balance
    creditCard.Pay(700) // decrease 700 from balance
    creditCard.Balance() // return balance (0 + 1000 - 700 = 300)
}
```