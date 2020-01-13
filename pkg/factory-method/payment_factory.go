package factory_method

// Producter provides a payment interface
type Producter interface {
	Pay(amount int) error
	Replenish(amount int) error
	Balance() int
}
