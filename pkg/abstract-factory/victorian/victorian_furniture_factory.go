package victorian

// abstractFactory provides victorianFurnitureFactory interface
type abstractFactory interface {
	CreateCoffeeTable() coffeeTable
	CreateSofa() sofa
	CreateChair() chair
}

// victorianFurnitureFactory implements victorian furniture factory
type victorianFurnitureFactory struct {
}

// CreateCoffeeTable returns new victorianCoffeeTable
func (v *victorianFurnitureFactory) CreateCoffeeTable() coffeeTable {
	return &victorianCoffeeTable{}
}

// CreateSofa returns new victorianSofa
func (v *victorianFurnitureFactory) CreateSofa() sofa {
	return &victorianSofa{}
}

// CreateChair returns new victorianChair
func (v *victorianFurnitureFactory) CreateChair() chair {
	return &victorianChair{}
}

// NewVictorianFurnitureFactory returns new victorianFurnitureFactory
func NewVictorianFurnitureFactory() abstractFactory {
	return &victorianFurnitureFactory{}
}
