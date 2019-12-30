package observer

// publisher interface
type publisher interface {
	Attach(observer observer)
	Unsubscribe(observer observer)
	SetTitleAndData(title, data string)
	Notify()
}

// observer provides a subscriber interface
type observer interface {
	update(title, data string)
	GetData() string
}

// newsPublisher implements the publisher interface
type newsPublisher struct {
	observers []observer
	title     string
	data      string
}

// subscriberPolitics implements observer interface
type subscriberPolitics struct {
	data string
}

// subscriberSport implements observer interface
type subscriberSport struct {
	data string
}

// Attach adds observer to newsPublisher
func (p *newsPublisher) Attach(observer observer) {
	p.observers = append(p.observers, observer)
}

// Unsubscribe removes observer from newsPublisher
func (p *newsPublisher) Unsubscribe(observer observer) {
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

// GetData returns subscriberPolitics data
func (s *subscriberPolitics) GetData() string {
	return s.data
}

// GetData returns subscriberSport data
func (s *subscriberSport) GetData() string {
	return s.data
}

// update sets new data if title matches with subscriberPolitics title
func (s *subscriberPolitics) update(title, data string) {
	if title == "Politics" {
		s.data = data
	}
}

// update sets new data if title matches with subscriberSport title
func (s *subscriberSport) update(title, data string) {
	if title == "Sport" {
		s.data = data
	}
}

// NewNewsPublisher returns new newsPublisher
func NewNewsPublisher() publisher {
	return &newsPublisher{}
}

// NewSubscriberPolitics returns new subscriberPolitics
func NewSubscriberPolitics() observer {
	return &subscriberPolitics{}
}

// NewSubscriberSport returns new subscriberSport
func NewSubscriberSport() observer {
	return &subscriberSport{}
}
