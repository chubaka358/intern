package factory_method

import (
	"errors"
	"fmt"
)

// creditCard implements credit card
type creditCard struct {
	amount int
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *creditCard) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("credit limit exceeded")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₽ со счета %T\n", amount, c)
	return nil
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *creditCard) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *creditCard) Balance() int {
	return c.amount
}

// NewCreditCard create and return new creditCard with creditCard.amount = 0
func NewCreditCard() *creditCard {
	return &creditCard{amount: 0}
}
