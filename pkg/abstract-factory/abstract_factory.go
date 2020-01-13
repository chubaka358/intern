package abstract_factory

// AbstractFactorer provides factory interface
type AbstractFactorer interface {
	CreateCoffeeTable() CoffeeTabler
	CreateSofa() Sofer
	CreateChair() Chairer
}
