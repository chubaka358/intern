package observer

// subscriberSport implements observer interface
type subscriberSport struct {
	data string
}

// GetData returns subscriberSport data
func (s *subscriberSport) GetData() string {
	return s.data
}

// update sets new data if title matches with subscriberSport title
func (s *subscriberSport) update(title, data string) {
	if title == "Sport" {
		s.data = data
	}
}

// NewSubscriberSport returns new subscriberSport
func NewSubscriberSport() Observer {
	return &subscriberSport{}
}
