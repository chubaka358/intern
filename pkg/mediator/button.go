package mediator

// button provides a button colleague
type button struct {
	mediator mediator
}

// SetMediator sets mediator
func (b *button) SetMediator(mediator mediator) {
	b.mediator = mediator
}

// SendData sends data
func (b *button) SendData() string {
	return b.mediator.Notify("button message")
}

// NewButton returns new button
func NewButton() *button {
	return &button{}
}
