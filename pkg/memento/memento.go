package memento

// memento implements memento
type memento struct {
	text        string
	cursorPos   int
	currentFont int
}

// getText returns memento's text
func (m *memento) getText() string {
	return m.text
}

// getCursorPos returns memento's cursorPos
func (m *memento) getCursorPos() int {
	return m.cursorPos
}

// getCurrentFont returns memento's currentFont
func (m *memento) getCurrentFont() int {
	return m.currentFont
}
