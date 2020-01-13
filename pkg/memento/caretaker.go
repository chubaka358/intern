package memento

// caretaker implements caretaker
type caretaker struct {
	memento []*memento
}

// AddMemento append memento to mementos list
func (c *caretaker) AddMemento(memento *memento) {
	c.memento = append(c.memento, memento)
}

// GetMemento returns memento with id=i from mementos list
func (c *caretaker) GetMemento(i int) *memento {
	return c.memento[i]
}

// GetMementos returns list of mementos
func (c *caretaker) GetMementos() []*memento {
	return c.memento
}

// NewCaretaker returns new caretaker
func NewCaretaker() *caretaker {
	return &caretaker{}
}
