package chain_of_responsibility

// secondHandler implements second handler
type secondHandler struct {
	Next Handler
}

// SendRequest implementation for secondHandler
func (h *secondHandler) SendRequest(data string) (result string) {
	if data == "data type 2" {
		result = "Using data type 2 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
}

// SetNextHandler set next handler for secondHandler
func (h *secondHandler) SetNextHandler(n Handler) Handler {
	h.Next = n
	return h
}

// NewSecondHandler return new secondHandler
func NewSecondHandler() Handler {
	return &secondHandler{}
}
