package memento

// editor implements editor
type editor struct {
	text        string
	cursorPos   int
	currentFont int
}

// SetText sets editor's text
func (e *editor) SetText(text string) {
	e.text = text
}

// GetText returns editor's text
func (e *editor) GetText() string {
	return e.text
}

// SetCursorPos sets editor's cursorPos
func (e *editor) SetCursorPos(cursorPos int) {
	e.cursorPos = cursorPos
}

// GetCursorPos returns editor's cursorPos
func (e *editor) GetCursorPos() int {
	return e.cursorPos
}

// SetCurrentFont sets editor's currentFont
func (e *editor) SetCurrentFont(currentFont int) {
	e.currentFont = currentFont
}

// GetCurrentFont returns editor's currentFont
func (e *editor) GetCurrentFont() int {
	return e.currentFont
}

// CreateMemento creates and returns memento with current editor state
func (e *editor) CreateMemento() *memento {
	return &memento{
		text:        e.text,
		cursorPos:   e.cursorPos,
		currentFont: e.currentFont,
	}
}

// SetMemento sets editor's state equal to memento state
func (e *editor) SetMemento(memento *memento) {
	e.text = memento.getText()
	e.cursorPos = memento.getCursorPos()
	e.currentFont = memento.getCurrentFont()
}

// NewEditor returns new editor
func NewEditor() *editor {
	return &editor{}
}
