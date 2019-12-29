package mediator

import "testing"

func TestMediator(t *testing.T) {
	button := &button{}
	textField := &textField{}
	checkbox := &checkbox{}
	Connect(button, textField, checkbox)

	gotTextField := textField.SendData()
	wantTextField := "data was filled"

	gotCheckbox := checkbox.SendData()
	wantCheckbox := "checkbox was checked"

	gotButton := button.SendData()
	wantButton := "form has been submitted"

	if gotTextField != wantTextField {
		t.Errorf("want %q, got %q", wantTextField, gotTextField)
	}
	if gotCheckbox != wantCheckbox {
		t.Errorf("want %q, got %q", wantCheckbox, gotCheckbox)
	}
	if gotButton != wantButton {
		t.Errorf("want %q, got %q", wantButton, gotButton)
	}

	wantSecondCheckbox := "already sent"
	wantSecondTextField := "already sent"
	gotSecondCheckbox := checkbox.SendData()
	gotSecondTextField := textField.SendData()
	if wantSecondCheckbox != gotSecondCheckbox {
		t.Errorf("want %q, got %q", wantSecondCheckbox, gotSecondCheckbox)
	}
	if wantSecondTextField != gotSecondTextField {
		t.Errorf("want %q, got %q", wantSecondTextField, gotSecondTextField)
	}
}
