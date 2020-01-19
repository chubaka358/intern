package victorian

// coffeeTable provides victorianCoffeeTable interface
type coffeeTable interface {
	Put() string
}

// victorianCoffeeTable implements victorian coffee table
type victorianCoffeeTable struct {
}

// Put puts item on the victorianCoffeeTable
func (v *victorianCoffeeTable) Put() string {
	return "Put item on the modernCoffeeTable"
}
