package visitor

// man implements Visitor interface
type man struct {
}

// VisitAirport implements visit to airport
func (m *man) VisitAirport(p *airport) string {
	return p.UseAirport()
}

// VisitSubway implements visit to subway
func (m *man) VisitSubway(p *subway) string {
	return p.UseSubway()
}

// VisitTeleport implements visit to teleport
func (m *man) VisitTeleport(p *teleport) string {
	return p.UseTeleport()
}

// NewMan creates and returns new man
func NewMan() *man {
	return &man{}
}
