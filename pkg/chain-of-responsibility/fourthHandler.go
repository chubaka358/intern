package chain_of_responsibility

// fourthHandler implements fourth handler
type fourthHandler struct {
	Next Handler
}

// SendRequest implementation for fourthHandler
func (h *fourthHandler) SendRequest(data string) (result string) {
	if data == "data type 4" {
		result = "Using data type 4 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
}

// SetNextHandler set next handler for fourthHandler
func (h *fourthHandler) SetNextHandler(n Handler) Handler {
	h.Next = n
	return h
}

// NewFourthHandler return new fourthHandler
func NewFourthHandler() Handler {
	return &fourthHandler{}
}
