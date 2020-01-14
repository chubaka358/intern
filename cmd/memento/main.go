package main

import (
	"github.com/jayhrat/intern/pkg/memento"
)

func main() {
	editor := memento.NewEditor()
	caretaker := memento.NewCaretaker()

	text := "Recorded in Athens and Thessaloniki, January to March 2018, while both cities were " +
		"being sold to airbnb."
	currentFont := 15
	cursorPos := 4
	editor.SetCurrentFont(currentFont)
	editor.SetCursorPos(cursorPos)
	editor.SetText(text)
	caretaker.AddMemento(editor.CreateMemento())
}
