package mediator

// coleaguer provides colleague interface
type colleaguer interface {
	SetMediator(mediator mediator)
	SendData() string
}
