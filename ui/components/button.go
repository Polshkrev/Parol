package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Custom button widget.
type Button struct {
	*widget.Button
	callback Callback
}

// Construct a new custom button widget.
// Returns a new button with a given text, icon, and callback.
func NewButton(text string, icon fyne.Resource, callback Callback) *Button {
	var button *Button = new(Button)
	button.callback = callback
	button.Button = &widget.Button{Text: text, Icon: icon, OnTapped: callback}
	button.ExtendBaseWidget(button)
	return button
}

// Custom tapped override.
func (button *Button) Tapped(_ *fyne.PointEvent) {
	if button.callback == nil {
		return
	}
	button.callback()
}

// Custom cursor override.
// Returns a pointer cursor.
func (button *Button) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
