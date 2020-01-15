package chain_of_responsibility

// Handler provides a handler interface
type Handler interface {
	SendRequest(data string) string
}
