package chain_of_responsibility

// firstHandler implements first handler
type firstHandler struct {
	next Handler
}

// SendRequest implementation for firstHandler
func (h *firstHandler) SendRequest(data string) string {
	if data == "data type 1" {
		return "Using data type 1 handler"
	}
	return h.next.SendRequest(data)
}

// NewFirstHandler return new firstHandler
func NewFirstHandler(next Handler) Handler {
	return &firstHandler{
		next: next,
	}
}
