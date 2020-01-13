package memento

// memento implements memento
type memento struct {
	text        string
	cursorPos   int
	currentFont int
}

// GetText returns memento's text
func (m *memento) getText() string {
	return m.text
}

// GetCursorPos returns memento's cursorPos
func (m *memento) getCursorPos() int {
	return m.cursorPos
}

// GetCurrentFont returns memento's currentFont
func (m *memento) getCurrentFont() int {
	return m.currentFont
}
