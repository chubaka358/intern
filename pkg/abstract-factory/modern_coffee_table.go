package abstract_factory

// modernCoffeeTable implements modern coffee table
type modernCoffeeTable struct {
}

// PutItem puts item on the modernCoffeeTable
func (m *modernCoffeeTable) PutItem() string {
	return "Put item on the modernCoffeeTable"
}
