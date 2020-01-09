package memento

// editorer provides editor interface
type editorer interface {
	SetText(text string)
	SetCursorPos(cursorPos int)
	SetCurrentFont(currentFont int)
	GetText() string
	GetCursorPos() int
	GetCurrentFont() int
	CreateMemento() *memento
	SetMemento(memento *memento)
}

// mementoer provides memento interface
type mementoer interface {
	getText() string
	getCursorPos() int
	getCurrentFont() int
}

// caretakerer provides caretaker interface
type caretakerer interface {
	AddMemento(memento *memento)
	GetMemento(i int) *memento
	GetMementos() []*memento
}

// editor implements editor
type editor struct {
	text        string
	cursorPos   int
	currentFont int
}

// memento implements memento
type memento struct {
	text        string
	cursorPos   int
	currentFont int
}

// caretaker implements caretaker
type caretaker struct {
	memento []*memento
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

// AddMemento append memento to mementos list
func (c *caretaker) AddMemento(memento *memento) {
	c.memento = append(c.memento, memento)
}

// GetMemento returns memento with id=i from mementos list
func (c *caretaker) GetMemento(i int) *memento {
	return c.memento[i]
}

// GetMementos returns list of mementos
func (c *caretaker) GetMementos() []*memento {
	return c.memento
}

// NewEditor returns new editor
func NewEditor() *editor {
	return &editor{}
}

// NewCaretaker returns new caretaker
func NewCaretaker() *caretaker {
	return &caretaker{}
}
