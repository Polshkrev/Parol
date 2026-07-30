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
func (button *Button) Tapped(*fyne.PointEvent) {
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

// Typed key event override.
func (button *Button) TypedKey(event *fyne.KeyEvent) {
	if event.Name != fyne.KeyReturn {
		button.Button.TypedKey(event)
	} else if button.callback == nil {
		return
	}
	button.callback()
}

// Make the default callback used in the button.
// Returns the default callback to use with the button.
func makeDefaultCallback(content string) Callback {
	return func() { fyne.CurrentApp().Clipboard().SetContent(content) }
}
