package visitor

// Place provides interface for place that the visitor should visit
type Place interface {
	Accept(v Visitor) string
}
