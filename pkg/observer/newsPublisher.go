package observer

// newsPublisher implements the publisher interface
type newsPublisher struct {
	observers []Observer
	title     string
	data      string
}

// Attach adds observer to newsPublisher
func (p *newsPublisher) Attach(observer Observer) {
	p.observers = append(p.observers, observer)
}

// Unsubscribe removes observer from newsPublisher
func (p *newsPublisher) Unsubscribe(observer Observer) {
	for i, o := range p.observers {
		if o == observer {
			p.observers[i] = p.observers[len(p.observers)-1]
			p.observers[len(p.observers)-1] = nil
			p.observers = p.observers[:len(p.observers)-1]
		}
	}
}

// SetTitleAndData sets new title and data
func (p *newsPublisher) SetTitleAndData(title, data string) {
	p.title = title
	p.data = data
}

// Notify sens notifications to subscribers
func (p *newsPublisher) Notify() {
	for _, observer := range p.observers {
		observer.update(p.title, p.data)
	}
}

// NewNewsPublisher returns new newsPublisher
func NewNewsPublisher() Publisher {
	return &newsPublisher{}
}
