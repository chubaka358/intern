package mediator

// mediator provides a mediator interface
type mediator interface {
	Notify(msg string) string
	Connect(button *button, textField *textField, checkbox *checkbox)
}
