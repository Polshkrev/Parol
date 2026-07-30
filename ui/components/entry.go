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

// Add the given callback to the given entry.
func addCallback(entry *Entry, callback Callback) {
	if !entry.submit || callback == nil {
		return
	}
	entry.OnSubmitted = func(string) {
		callback()
	}
}

// Construct a new custom entry widget.
// Returns a new custom entry widget.
func NewEntry(text string, password bool, submit bool, validationFunction fyne.StringValidator, submitCallback Callback) *Entry {
	var entry *Entry = new(Entry)
	entry.text = text
	entry.password = password
	entry.submit = submit
	entry.validationFunction = validationFunction
	addEntryBase(entry)
	addCallback(entry, submitCallback)
	entry.ExtendBaseWidget(entry)
	return entry
}

// Obtain the value of the entry.
// Returns the text of the entry.
func (entry *Entry) Value() string {
	return entry.Text
}
