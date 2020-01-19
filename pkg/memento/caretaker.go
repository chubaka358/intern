package memento

// Caretaker provides caretaker interface
type Caretaker interface {
	AddMemento(memento Memento)
	GetMemento(i int) Memento
	GetMementos() []Memento
}

// caretaker implements caretaker
type caretaker struct {
	memento []Memento
}

// AddMemento append memento to mementos list
func (c *caretaker) AddMemento(memento Memento) {
	c.memento = append(c.memento, memento)
}

// GetMemento returns memento with id=i from mementos list
func (c *caretaker) GetMemento(i int) Memento {
	return c.memento[i]
}

// GetMementos returns list of mementos
func (c *caretaker) GetMementos() []Memento {
	return c.memento
}

// NewCaretaker returns new caretaker
func NewCaretaker() Caretaker {
	return &caretaker{}
}
