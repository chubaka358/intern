package mediator

// actionMediator implements mediator
type actionMediator struct {
	*button
	*textField
	*checkbox
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

// Connect connects all colleagues
func (a *actionMediator) Connect(button *button, textField *textField, checkbox *checkbox) {
	a.button = button
	a.textField = textField
	a.checkbox = checkbox

	a.checkbox.SetMediator(a)
	a.textField.SetMediator(a)
	a.button.SetMediator(a)
}

// NewActionMediator returns new actionMediator
func NewActionMediator() *actionMediator {
	return &actionMediator{}
}
