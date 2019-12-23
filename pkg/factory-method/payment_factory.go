package factory_method

import (
	"errors"
	"fmt"
)

// Creater provides a factory interface
type Creater interface {
	CreatePayment(string) (Producter, error)
}

// Cash implements hard cash
type Cash struct {
	amount int
}

// CreditCard implements credit card
type CreditCard struct {
	amount int
}

// Bitcoin implements bitcoin
type Bitcoin struct {
	amount int
}

// Producter provides a payment interface
type Producter interface {
	Pay(amount int) error
	Replenish(amount int) error
	Balance() int
}

// PaymentFactory implements Creater intrface
type PaymentFactory struct {
}

// NewFactory is the PaymentFactory constructor
func NewFactory() Creater {
	return &PaymentFactory{}
}

// CreatePayment creates Producter
func (p *PaymentFactory) CreatePayment(action string) (Producter, error) {
	var payment Producter
	switch action {
	case "Cash":
		return &Cash{amount: 0}, nil
	case "CreditCard":
		return &CreditCard{amount: 0}, nil
	case "Bitcoin":
		return &Bitcoin{amount: 0}, nil
	}
	return payment, errors.New("unknown payment")
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *Cash) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("not enough money")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₽ со счета %T\n", amount, c)
	return nil
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *Cash) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *Cash) Balance() int {
	return c.amount
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *CreditCard) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("credit limit exceeded")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₽ со счета %T\n", amount, c)
	return nil
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *CreditCard) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *CreditCard) Balance() int {
	return c.amount
}

// Pay allows you to write off money from the account
// Throws an error if there is not enough money
func (c *Bitcoin) Pay(amount int) error {
	if c.amount-amount < 0 {
		return errors.New("the transaction has not been approved")
	}
	c.amount -= amount
	fmt.Printf("Списано %v₿ со счета %T\n", amount, c)
	return nil
}

// Replenish allows you to replenish the balance
// Throws an error if we try to replenish by less than or equal to zero
func (c *Bitcoin) Replenish(amount int) error {
	if amount <= 0 {
		return errors.New("expected replenishment of the balance in the amount greater than zero")
	}
	c.amount += amount
	return nil
}

// Balance returns account balance
func (c *Bitcoin) Balance() int {
	return c.amount
}
