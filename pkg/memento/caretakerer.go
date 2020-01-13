package memento

// Caretakerer provides caretaker interface
type Caretakerer interface {
	AddMemento(memento *memento)
	GetMemento(i int) *memento
	GetMementos() []*memento
}
