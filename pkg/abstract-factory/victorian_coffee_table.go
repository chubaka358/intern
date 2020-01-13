package abstract_factory

// victorianCoffeeTable implements victorian coffee table
type victorianCoffeeTable struct {
}

// PutItem puts item on the victorianCoffeeTable
func (v *victorianCoffeeTable) PutItem() string {
	return "Put item on the modernCoffeeTable"
}
