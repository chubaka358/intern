package chain_of_responsibility

// Handler provides a handler interface
type Handler interface {
	SendRequest(requestType string, data string) (result string)
}

// GetHandler implements handler "GET"
type GetHandler struct {
	Next Handler
}

// SendRequest implementation for GetHandler
func (h *GetHandler) SendRequest(requestType string, data string) (result string) {
	if requestType == "GET" {
		result = "Using GET method"
	} else if h.Next != nil {
		result = h.Next.SendRequest(requestType, data)
	}
	return
}

// PostHandler implements handler "POST"
type PostHandler struct {
	Next Handler
}

// SendRequest implementation for PostHandler
func (h *PostHandler) SendRequest(requestType string, data string) (result string) {
	if requestType == "POST" {
		result = "Using POST method"
	} else if h.Next != nil {
		result = h.Next.SendRequest(requestType, data)
	}
	return
}

// PutHandler implements handler "PUT"
type PutHandler struct {
	Next Handler
}

// SendRequest implementation for PutHandler
func (h *PutHandler) SendRequest(requestType string, data string) (result string) {
	if requestType == "PUT" {
		result = "Using PUT method"
	} else if h.Next != nil {
		result = h.Next.SendRequest(requestType, data)
	}
	return
}

// PutHandler implements handler "PUT"
type DeleteHandler struct {
	Next Handler
}

// SendRequest implementation for SendHandler
func (h *DeleteHandler) SendRequest(requestType string, data string) (result string) {
	if requestType == "DELETE" {
		result = "Using DELETE method"
	} else if h.Next != nil {
		result = h.Next.SendRequest(requestType, data)
	}
	return
}
