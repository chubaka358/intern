package memento

// Mementoer provides memento interface
type Mementoer interface {
	getText() string
	getCursorPos() int
	getCurrentFont() int
}
