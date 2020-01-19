package modern

// abstractFactory provides modernFurnitureFactory interface
type abstractFactory interface {
	CreateCoffeeTable() coffeeTable
	CreateSofa() sofa
	CreateChair() chair
}

// modernFurnitureFactory implements modern furniture factory
type modernFurnitureFactory struct {
}

// CreateCoffeeTable returns new modernCoffeeTable
func (m *modernFurnitureFactory) CreateCoffeeTable() coffeeTable {
	return &modernCoffeeTable{}
}

// CreateSofa returns new modernSofa
func (m *modernFurnitureFactory) CreateSofa() sofa {
	return &modernSofa{}
}

// CreateChair returns new modernChair
func (m *modernFurnitureFactory) CreateChair() chair {
	return &modernChair{}
}

// NewModernFurnitureFactory returns new modernFurnitureFactory
func NewModernFurnitureFactory() abstractFactory {
	return &modernFurnitureFactory{}
}
