package chain_of_responsibility

// Handler provides a handler interface
type handler interface {
	SendRequest(data string) (result string)
	SetNextHandler(n handler) handler
}

// firstHandler implements first handler
type firstHandler struct {
	Next handler
}

// secondHandler implements second handler
type secondHandler struct {
	Next handler
}

// thirdHandler implements third handler
type thirdHandler struct {
	Next handler
}

// fourthHandler implements fourth handler
type fourthHandler struct {
	Next handler
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

// SendRequest implementation for secondHandler
func (h *secondHandler) SendRequest(data string) (result string) {
	if data == "data type 2" {
		result = "Using data type 2 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
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

// SendRequest implementation for fourthHandler
func (h *fourthHandler) SendRequest(data string) (result string) {
	if data == "data type 4" {
		result = "Using data type 4 handler"
	} else if h.Next != nil {
		result = h.Next.SendRequest(data)
	}
	return
}

// SetNextHandler set next handler for firstHandler
func (h *firstHandler) SetNextHandler(n handler) handler {
	h.Next = n
	return h
}

// SetNextHandler set next handler for secondHandler
func (h *secondHandler) SetNextHandler(n handler) handler {
	h.Next = n
	return h
}

// SetNextHandler set next handler for thirdHandler
func (h *thirdHandler) SetNextHandler(n handler) handler {
	h.Next = n
	return h
}

// SetNextHandler set next handler for fourthHandler
func (h *fourthHandler) SetNextHandler(n handler) handler {
	h.Next = n
	return h
}

// NewFirstHandler return new firstHandler
func NewFirstHandler() handler {
	return &firstHandler{}
}

// NewSecondHandler return new secondHandler
func NewSecondHandler() handler {
	return &secondHandler{}
}

// NewThirdHandler return new thirdHandler
func NewThirdHandler() handler {
	return &thirdHandler{}
}

// NewFourthHandler return new fourthHandler
func NewFourthHandler() handler {
	return &fourthHandler{}
}
