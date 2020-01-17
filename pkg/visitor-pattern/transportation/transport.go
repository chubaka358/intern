package transportation

// Transport provides interface for transportation that the visitor should visit
type Transport interface {
	Accept(v visitor) string
	Use() string
}
