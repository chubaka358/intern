package visitor

type Visitor interface {
	VisitAirport(p *airport) string
	VisitSubway(p *subway) string
	VisitTeleport(p *teleport) string
}
