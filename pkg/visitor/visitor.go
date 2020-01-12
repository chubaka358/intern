package visitor

// Visitor provides a visitor interface
type Visitor interface {
	VisitAirport(p *airport) string
	VisitSubway(p *subway) string
	VisitTeleport(p *teleport) string
}
