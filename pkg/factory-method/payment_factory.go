package factory_method

import (
	"errors"
	"fmt"
)

// Cash implements hard cash
type cash struct {
	amount int
}

// CreditCard implements credit card
type creditCard struct {
	amount int
}

// Bitcoin implements bitcoin
type bitcoin struct {
	amount int
}

// Producter provides a payment interface
type producter interface {
	Pay(amount int) error
	Replenish(amount int) error
	Balance() int
}

func NewCash() producter {
	return &cash{amount: 0}
}

func NewCreditCard() producter {
	return &creditCard{amount: 0}
}

func NewBitcoin() producter {
	return &bitcoin{amount: 0}
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

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *cash) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *cash) Balance() int {
	return c.amount
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
