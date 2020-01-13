package chain_of_responsibility

// thirdHandler implements third handler
type thirdHandler struct {
	Next Handler
}

// SendRequest implementation for thirdHandler
func (h *thirdHandler) SendRequest(data string) (result string) {
	if data == "data type 3" {
		result = "Using data type 3 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
}

// SetNextHandler set next handler for thirdHandler
func (h *thirdHandler) SetNextHandler(n Handler) Handler {
	h.Next = n
	return h
}

// NewThirdHandler return new thirdHandler
func NewThirdHandler() Handler {
	return &thirdHandler{}
}
