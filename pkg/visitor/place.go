package visitor

type Place interface {
	Accept(v Visitor) string
}
