package main

import (
	"fmt"
	"github.com/chubaka358/intern/pkg/factory-method"
)

func main() {
	factory := factory_method.NewFactory()

	if cash, err := factory.CreatePayment("Cash"); err == nil {
		payErr := cash.Pay(1000)
		if payErr != nil {
			fmt.Println(payErr)
		}
	} else {
		fmt.Println(err)
	}

	if creditCard, err := factory.CreatePayment("CreditCard"); err == nil {
		payErr := creditCard.Pay(1001)
		if payErr != nil {
			fmt.Println(payErr)
		}
	} else {
		fmt.Println(err)
	}

	if bitcoin, err := factory.CreatePayment("Bitcoin"); err == nil {
		payErr := bitcoin.Pay(3)
		if payErr != nil {
			fmt.Println(payErr)
		}
	} else {
		fmt.Println(err)
	}
}
