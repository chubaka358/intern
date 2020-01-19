package transportation

import "github.com/jayhrat/intern/pkg/visitor-pattern/visitor"

// Transport provides interface for transportation that the visitor should visit
type Transport interface {
	Accept(v visitor.Visitor) string
	Use() string
}
