package visitor

type airport struct {
}

func (a *airport) UseAirport() string {
	return "using airport..."
}

func (a *airport) Accept(v Visitor) string {
	return v.VisitAirport(a)
}

func NewAirport() *airport {
	return &airport{}
}
