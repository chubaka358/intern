package modern

// coffeeTable provides modernCoffeeTable interface
type coffeeTable interface {
	Put() string
}

// modernCoffeeTable implements modern coffee table
type modernCoffeeTable struct {
}

// Put puts item on the modernCoffeeTable
func (m *modernCoffeeTable) Put() string {
	return "Put item on the modernCoffeeTable"
}
