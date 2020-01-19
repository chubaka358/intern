package abstract

// AbstractFactory provides factory interface
type AbstractFactory interface {
	CreateCoffeeTable() CoffeeTable
	CreateSofa() Sofa
	CreateChair() Chair
}
