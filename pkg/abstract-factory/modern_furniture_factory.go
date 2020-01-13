package abstract_factory

// modernFurnitureFactory implements modern furniture factory
type modernFurnitureFactory struct {
}

// CreateCoffeeTable returns new modernCoffeeTable
func (m *modernFurnitureFactory) CreateCoffeeTable() CoffeeTabler {
	return &modernCoffeeTable{}
}

// CreateSofa returns new modernSofa
func (m *modernFurnitureFactory) CreateSofa() Sofer {
	return &modernSofa{}
}

// CreateChair returns new modernChair
func (m *modernFurnitureFactory) CreateChair() Chairer {
	return &modernChair{}
}

// NewModernFurnitureFactory returns new modernFurnitureFactory
func NewModernFurnitureFactory() AbstractFactorer {
	return &modernFurnitureFactory{}
}
