package chain_of_responsibility

// terminator implements terminator
type terminator struct {
}

// SendRequest implementation for terminator
func (t *terminator) SendRequest(data string) string {
	return ""
}

// NewTerminator return new terminator
func NewTerminator() Handler {
	return &terminator{}
}
