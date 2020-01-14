package mediator

// actionMediator implements mediator
type actionMediator struct {
	*button
	*textField
	*checkbox
}

// Notify implementation
func (a *actionMediator) Notify(msg string) string {
	var answer string
	switch msg {
	case "button message":
		a.textField.sent = true
		a.checkbox.sent = true
		answer = "form has been submitted"
	case "textField message":
		if a.textField.sent {
			return "already sent"
		}
		a.textField.dataFilled = true
		answer = "data was filled"
	case "checkbox message":
		if a.textField.sent {
			return "already sent"
		}
		a.checkbox.state = true
		answer = "checkbox was checked"
	default:
		answer = ""
	}
	return answer
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
