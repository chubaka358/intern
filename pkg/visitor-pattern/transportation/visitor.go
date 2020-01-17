package transportation

// visitor provides a visitor interface
type visitor interface {
	Visit(p Transport) string
}
