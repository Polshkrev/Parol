package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Custom entry widget.
type Entry struct {
	*widget.Entry
	text               string
	password           bool
	submit             bool
	validationFunction fyne.StringValidator
}

// Add the custom base of the given entry.
func addEntryBase(entry *Entry) {
	if entry.password {
		entry.Entry = widget.NewPasswordEntry()
	} else {
		entry.Entry = widget.NewEntry()
	}
	entry.Validator = entry.validationFunction
	entry.SetPlaceHolder(entry.text)
}

// Construct a new custom entry widget.
// Returns a new custom entry widget.
func NewEntry(text string, password bool, submit bool, validationFunction fyne.StringValidator) *Entry {
	var entry *Entry = new(Entry)
	entry.text = text
	entry.password = password
	entry.submit = submit
	entry.validationFunction = validationFunction
	addEntryBase(entry)
	entry.ExtendBaseWidget(entry.Entry)
	return entry
}

// Obtain the value of the entry.
// Returns the text of the entry.
func (entry *Entry) Value() string {
	return entry.Text
}
