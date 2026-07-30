package components

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Custom header widget.
type Header struct {
	label  *widget.Label
	parent *fyne.Container
	button *Button
	width  uint8
	height uint8
}

// Construct a new header based on its parts.
// Returns  a new header based on its parts.
func NewHeader(label *widget.Label, button *Button, width, height uint8) *Header {
	var header *Header = new(Header)
	header.label = label
	header.button = button
	header.width = width
	header.height = height
	return header
}

// Set the parent of the header.
func (header *Header) SetParent(parent *fyne.Container) {
	header.parent = parent
}

// Obtain the parent of the header.
// Returns the [fyne.Container] of the header.
func (header Header) Parent() *fyne.Container {
	return header.parent
}

// Default callback for searching for a card.
func searchCallback(value string, cardParent *fyne.Container) {
	var i int
	for i = range cardParent.Objects {
		var object *Card
		var ok bool
		object, ok = cardParent.Objects[i].(*Card)
		if !ok {
			continue
		} else if strings.Contains(object.Key(), strings.ToLower(value)) {
			object.Show()
		} else {
			object.Hide()
		}
	}
	cardParent.Refresh()
}

// Paint the search bar.
// Returns the canvas object containing the search bar.
func paintSearch(cardParent *fyne.Container) fyne.CanvasObject {
	var search *Entry = NewEntry(lang.L("gui.labels.prompts.search"), false, false, nil, nil)
	search.Wrapping = fyne.TextWrapOff
	search.Scroll = container.ScrollHorizontalOnly
	search.OnChanged = func(value string) { searchCallback(value, cardParent) }
	var item *Item = NewItem(search)
	var searchContainer *fyne.Container = container.NewCenter(container.NewGridWrap(fyne.NewSize(250, 35), item.Object()))
	return searchContainer
}

// Paint the header.
func (header *Header) Paint(cardParent *fyne.Container) {
	var frame *fyne.Container = container.New(layout.NewGridWrapLayout(fyne.NewSize(float32(header.width), float32(header.height))), header.button)
	header.SetParent((container.NewBorder(nil, widget.NewSeparator(), container.NewCenter(header.label), container.NewPadded(frame), paintSearch(cardParent))))
}
