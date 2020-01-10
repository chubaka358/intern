package visitor

// man implements Visitor interface
type man struct {
}

func (m *man) VisitAirport(p *airport) string {
	return p.UseAirport()
}

func (m *man) VisitSubway(p *subway) string {
	return p.UseSubway()
}

func (m *man) VisitTeleport(p *teleport) string {
	return p.UseTeleport()
}

func NewMan() *man {
	return &man{}
}
