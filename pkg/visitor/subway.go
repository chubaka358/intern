package visitor

type subway struct {
}

func (s *subway) UseSubway() string {
	return "using subway..."
}

func (s *subway) Accept(v Visitor) string {
	return v.VisitSubway(s)
}

func NewSubway() *subway {
	return &subway{}
}
