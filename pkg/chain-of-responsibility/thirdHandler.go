package chain_of_responsibility

// thirdHandler implements third handler
type thirdHandler struct {
	next Handler
}

// SendRequest implementation for thirdHandler
func (h *thirdHandler) SendRequest(data string) string {
	if data == "data type 3" {
		return "Using data type 3 handler"
	}
	return h.next.SendRequest(data)
}

// NewThirdHandler return new thirdHandler
func NewThirdHandler(next Handler) Handler {
	return &thirdHandler{
		next: next,
	}
}
