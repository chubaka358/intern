package transportation

// Transport provides interface for transportation that the visitor-pattern should visit
type Transport interface {
	Accept(v visitor) string
	Use() string
}
