package memento

import "testing"

func TestMemento(t *testing.T) {
	editor := NewEditor()
	caretaker := NewCaretaker()

	text := "Recorded in Athens and Thessaloniki, January to March 2018, while both cities were " +
		"being sold to airbnb."
	currentFont := 15
	cursorPos := 4

	editor.SetCurrentFont(currentFont)
	editor.SetCursorPos(cursorPos)
	editor.SetText(text)
	caretaker.AddMemento(editor.CreateMemento())
	editor.SetCurrentFont(7)
	editor.SetCursorPos(2)
	editor.SetText("по телу солдата ползают муравьи " +
		"заползают в уши ноздри глаза рот " +
		"муравьи вгрызаются в череп " +
		"и стройным маршем идут жрать " +
		"муравьи едят мозг " +
		"едят солдатское тело")
	editor.SetMemento(caretaker.GetMemento(0))
	if got := editor.GetText(); got != text {
		t.Errorf("expected %q, got %q", text, got)
	}
	if got := editor.GetCurrentFont(); got != currentFont {
		t.Errorf("expected %q, got %q", text, got)
	}
	if got := editor.GetCursorPos(); got != cursorPos {
		t.Errorf("expected %q, got %q", text, got)
	}
}
