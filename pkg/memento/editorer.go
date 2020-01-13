package memento

// Editorer provides editor interface
type Editorer interface {
	SetText(text string)
	SetCursorPos(cursorPos int)
	SetCurrentFont(currentFont int)
	GetText() string
	GetCursorPos() int
	GetCurrentFont() int
	CreateMemento() *memento
	SetMemento(memento *memento)
}
