package mediator

// textField provides a textField colleague
type textField struct {
	mediator   mediator
	dataFilled bool
	sent       bool
}

// SetMediator sets mediator
func (t *textField) SetMediator(mediator mediator) {
	t.mediator = mediator
}

// SendData sends data
func (t *textField) SendData() string {
	return t.mediator.Notify("textField message")
}

// NewTextField returns new textField
func NewTextField() *textField {
	return &textField{}
}
