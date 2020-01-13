package chain_of_responsibility

// firstHandler implements first handler
type firstHandler struct {
	Next Handler
}

// SendRequest implementation for firstHandler
func (h *firstHandler) SendRequest(data string) (result string) {
	if data == "data type 1" {
		result = "Using data type 1 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
}

// SetNextHandler set next handler for firstHandler
func (h *firstHandler) SetNextHandler(n Handler) Handler {
	h.Next = n
	return h
}

// NewFirstHandler return new firstHandler
func NewFirstHandler() Handler {
	return &firstHandler{}
}
