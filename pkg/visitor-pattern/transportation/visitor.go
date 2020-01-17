package transportation

// Visitor provides a visitor-pattern interface
type visitor interface {
	Visit(p Transport) string
}
