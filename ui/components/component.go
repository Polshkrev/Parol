package components

import "fyne.io/fyne/v2"

// Type alias for a component callback.
type Callback func()

// Standardization of a gui component.
type Component interface {
	// Paint a component to a given parent container.
	Paint(parent *fyne.Container)
}
