package chain_of_responsibility

// secondHandler implements second handler
type secondHandler struct {
	next Handler
}

// SendRequest implementation for secondHandler
func (h *secondHandler) SendRequest(data string) string {
	if data == "data type 2" {
		return "Using data type 2 handler"
	}
	return h.next.SendRequest(data)
}

// NewSecondHandler return new secondHandler
func NewSecondHandler(next Handler) Handler {
	return &secondHandler{
		next: next,
	}
}
