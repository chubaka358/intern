package chain_of_responsibility

// fourthHandler implements fourth handler
type fourthHandler struct {
	next Handler
}

// SendRequest implementation for fourthHandler
func (h *fourthHandler) SendRequest(data string) string {
	if data == "data type 4" {
		return "Using data type 4 handler"
	}
	return h.next.SendRequest(data)
}

// NewFourthHandler return new fourthHandler
func NewFourthHandler(next Handler) Handler {
	return &fourthHandler{
		next: next,
	}
}
