package abstract_factory

// victorianFurnitureFactory implements victorian furniture factory
type victorianFurnitureFactory struct {
}

// CreateCoffeeTable returns new victorianCoffeeTable
func (v *victorianFurnitureFactory) CreateCoffeeTable() CoffeeTabler {
	return &victorianCoffeeTable{}
}

// CreateSofa returns new victorianSofa
func (v *victorianFurnitureFactory) CreateSofa() Sofer {
	return &victorianSofa{}
}

// CreateChair returns new victorianChair
func (v *victorianFurnitureFactory) CreateChair() Chairer {
	return &victorianChair{}
}

// NewVictorianFurnitureFactory returns new victorianFurnitureFactory
func NewVictorianFurnitureFactory() AbstractFactorer {
	return &victorianFurnitureFactory{}
}
