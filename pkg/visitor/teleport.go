package visitor

type teleport struct {
}

func (t *teleport) UseTeleport() string {
	return "using teleport..."
}

func (t *teleport) Accept(v Visitor) string {
	return v.VisitTeleport(t)
}

func NewTeleport() *teleport {
	return &teleport{}
}
