package factory_method

import (
	"errors"
	"fmt"
)

// bitcoin implements bitcoin
type bitcoin struct {
	amount int
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *bitcoin) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("the transaction has not been approved")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₿ со счета %T\n", amount, c)
	return nil
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *bitcoin) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *bitcoin) Balance() int {
	return c.amount
}

// NewBitcoin create and return new bitcoin with bitcoin.amount = 0
func NewBitcoin() Producter {
	return &bitcoin{
		amount: 0,
	}
}
