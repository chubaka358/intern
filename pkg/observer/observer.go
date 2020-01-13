package observer

// observer provides a subscriber interface
type Observer interface {
	update(title, data string)
	GetData() string
}
