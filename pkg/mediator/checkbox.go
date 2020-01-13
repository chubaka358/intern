package mediator

// checkbox provides a checkbox colleague
type checkbox struct {
	mediator mediator
	state    bool
	sent     bool
}

// SetMediator sets mediator
func (c *checkbox) SetMediator(mediator mediator) {
	c.mediator = mediator
}

// SendData sends data
func (c *checkbox) SendData() string {
	return c.mediator.Notify("checkbox message")
}

// NewCheckbox return new checkbox
func NewCheckbox() *checkbox {
	return &checkbox{}
}
