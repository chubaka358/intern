package observer

// subscriberPolitics implements observer interface
type subscriberPolitics struct {
	data string
}

// GetData returns subscriberPolitics data
func (s *subscriberPolitics) GetData() string {
	return s.data
}

// update sets new data if title matches with subscriberPolitics title
func (s *subscriberPolitics) update(title, data string) {
	if title == "Politics" {
		s.data = data
	}
}

// NewSubscriberPolitics returns new subscriberPolitics
func NewSubscriberPolitics() Observer {
	return &subscriberPolitics{}
}
