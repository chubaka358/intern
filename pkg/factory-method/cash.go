package factory_method

import (
	"errors"
	"fmt"
)

// cash implements hard cash
type cash struct {
	amount int
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *cash) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("not enough money")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₽ со счета %T\n", amount, c)
	return nil
}

// Balance returns account balance
func (c *cash) Balance() int {
	return c.amount
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *cash) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// NewCash create and return new cash with cash.amount = 0
func NewCash() *cash {
	return &cash{
		amount: 0,
	}
}
