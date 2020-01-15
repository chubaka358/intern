package memento

// Caretakerer provides caretaker interface
type Caretakerer interface {
	AddMemento(memento Mementoer)
	GetMemento(i int) Mementoer
	GetMementos() []Mementoer
}
