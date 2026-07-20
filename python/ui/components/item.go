package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Custom form item.
type Item struct {
	*widget.FormItem
	entry *Entry
}

// Construct a new custom form item from a given entry.
// Returns a new custom form item from a given entry.
func NewItem(entry *Entry) *Item {
	var item *Item = new(Item)
	item.entry = entry
	item.FormItem = widget.NewFormItem(entry.text, entry.Entry)
	return item
}

// Obtain the base widget of the form item.
// Returns the base widget of the form item.
func (item *Item) Base() *widget.FormItem {
	return item.FormItem
}

// Obtain the base canvas object of the form item.
// Returns the base canvas object of the form item.
func (item *Item) Object() fyne.CanvasObject {
	return item.Widget
}

// Obtain the underlying entry of the form item.
// Returns the underlying entry of the form item.
func (item *Item) Entry() *Entry {
	return item.entry
}

// Obtain the value of the form item.
// Returns the value of the form item.
func (item *Item) Value() string {
	return item.entry.Value()
}
