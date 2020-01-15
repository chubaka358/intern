package memento

// caretaker implements caretaker
type caretaker struct {
	memento []Mementoer
}

// AddMemento append memento to mementos list
func (c *caretaker) AddMemento(memento Mementoer) {
	c.memento = append(c.memento, memento)
}

// GetMemento returns memento with id=i from mementos list
func (c *caretaker) GetMemento(i int) Mementoer {
	return c.memento[i]
}

// GetMementos returns list of mementos
func (c *caretaker) GetMementos() []Mementoer {
	return c.memento
}

// NewCaretaker returns new caretaker
func NewCaretaker() Caretakerer {
	return &caretaker{}
}
