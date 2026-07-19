package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

// Paint the header.
func (header *Header) Paint(*fyne.Container) {
	var frame *fyne.Container = container.New(layout.NewGridWrapLayout(fyne.NewSize(float32(header.width), float32(header.height))), header.button)
	header.SetParent((container.NewBorder(nil, widget.NewSeparator(), container.NewCenter(header.label), container.NewPadded(frame))))
}
