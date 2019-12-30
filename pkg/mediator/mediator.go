package mediator

// mediator provides a mediator interface
type mediator interface {
	Notify(msg string) string
}

// actionMediator implements mediator
type actionMediator struct {
	*button
	*textField
	*checkbox
}

// button provides a button colleague
type button struct {
	mediator mediator
}

// textField provides a textField colleague
type textField struct {
	mediator   mediator
	dataFilled bool
	sent       bool
}

// checkbox provides a checkbox colleague
type checkbox struct {
	mediator mediator
	state    bool
	sent     bool
}

// Notify implementation
func (a *actionMediator) Notify(msg string) string {
	if msg == "button message" {
		a.textField.sent = true
		a.checkbox.sent = true
		return "form has been submitted"
	} else if msg == "textField message" {
		if a.textField.sent {
			return "already sent"
		}
		a.textField.dataFilled = true
		return "data was filled"
	} else if msg == "checkbox message" {
		if a.textField.sent {
			return "already sent"
		}
		a.checkbox.state = true
		return "checkbox was checked"
	}
	return ""
}

// SetMediator sets mediator
func (b *button) SetMediator(mediator mediator) {
	b.mediator = mediator
}

// SendData sends data
func (b *button) SendData() string {
	return b.mediator.Notify("button message")
}

// SetMediator sets mediator
func (t *textField) SetMediator(mediator mediator) {
	t.mediator = mediator
}

// SendData sends data
func (t *textField) SendData() string {
	return t.mediator.Notify("textField message")
}

// SetMediator sets mediator
func (c *checkbox) SetMediator(mediator mediator) {
	c.mediator = mediator
}

// SendData sends data
func (c *checkbox) SendData() string {
	return c.mediator.Notify("checkbox message")
}

// Connect connects all colleagues
func Connect(button *button, textField *textField, checkbox *checkbox) {
	mediator := &actionMediator{
		button:    button,
		textField: textField,
		checkbox:  checkbox,
	}

	mediator.checkbox.SetMediator(mediator)
	mediator.textField.SetMediator(mediator)
	mediator.button.SetMediator(mediator)
}

// NewButton returns new button
func NewButton() *button {
	return &button{}
}

// NewTextField returns new textField
func NewTextField() *textField {
	return &textField{}
}

// NewCheckbox return new checkbox
func NewCheckbox() *checkbox {
	return &checkbox{}
}
