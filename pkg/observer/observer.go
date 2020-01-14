package observer

// Observer provides a subscriber interface
type Observer interface {
	update(title, data string)
	GetData() string
}
