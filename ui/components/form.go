package components

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

const (
	enCode string = "en" // Code to check if the system locale is `english`.
)

// Type alias for a form callback.
type FormCallback func(valid bool)

// Custom form dialog widget.
type Form struct {
	*dialog.CustomDialog
	parent         *fyne.Window
	title          string
	submitText     string
	cancelText     string
	width          uint16
	height         uint16
	submitCallback Callback
	cancelCallback Callback
	items          collections.Collection[*Item]
}

// Construct a new form widget.
// Returns a new custom form widget based on its given parts.
func NewForm(title, submitText, cancelText string, width, height uint16, submitCallback, cancelCallback Callback) *Form {
	var form *Form = new(Form)
	form.title = title
	form.submitText = submitText
	form.cancelText = cancelText
	form.width = width
	form.height = height
	form.submitCallback = submitCallback
	form.cancelCallback = cancelCallback
	form.items = collections.NewArray[*Item]()
	return form
}

// Append a form item to the form.
func (form *Form) Append(item *Item) {
	form.items.Append(item)
}

// Obtain the parent of the form.
// Returns the base window of the form.
func (form *Form) Parent() *fyne.Window {
	return form.parent
}

// Obtain the submit callback for the form.
// Returns the callback that triggers when the form is submitted.
func (form *Form) SubmitCallback() Callback {
	return form.submitCallback
}

// Obtain each of the form item data.
// Returns the data of each of the items stored within the form.
func (form *Form) Collect() []*widget.FormItem {
	var result []*widget.FormItem = make([]*widget.FormItem, 0)
	var i gopolutils.Size
	for i = range collections.Enumerate(form.items) {
		var item **Item = gopolutils.Must(form.items.At(i))
		result = append(result, (*item).Base())
	}
	return result
}

// Obtain the internal item widgets of the form.
// Returns a [collections.View] of [Item]s.
func (form *Form) Items() collections.View[*Item] {
	return form.items
}

// Returns each of the base canvas objects for the form.
// Returns a slice of [fyne.CanvasObject]s of the form.
func (form *Form) Objects() []fyne.CanvasObject {
	var result []fyne.CanvasObject = make([]fyne.CanvasObject, 0)
	var i gopolutils.Size
	for i = range collections.Enumerate(form.items) {
		var item **Item = gopolutils.Must(form.items.At(i))
		result = append(result, (*item).Object())
	}
	return result
}

// Paint a form item.
func (form *Form) Paint(*fyne.Container) {
	var submitButton *Button = NewButton(form.submitText, theme.ConfirmIcon(), form.submitCallback)
	var cancelButton *Button = NewButton(form.cancelText, theme.CancelIcon(), form.cancelCallback)
	submitButton.Importance = widget.HighImportance
	var box *fyne.Container
	if !strings.Contains(lang.SystemLocale().String(), enCode) {
		box = container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), container.NewGridWrap(fyne.NewSize(100, 40), cancelButton), container.NewGridWrap(fyne.NewSize(100, 40), submitButton))
	} else {
		box = container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), container.NewGridWrap(fyne.NewSize(85, 40), cancelButton), container.NewGridWrap(fyne.NewSize(75, 40), submitButton))
	}
	var buttons *fyne.Container = container.NewCenter(box)
	var entries *fyne.Container = container.NewVBox(form.Objects()...)
	var object *fyne.Container = container.NewVBox(entries, buttons)
	form.CustomDialog = dialog.NewCustomWithoutButtons(form.title, object, *form.Parent())
	var formSize fyne.Size = fyne.NewSize(float32(form.width), float32(form.height))
	(*form.parent).SetFixedSize(true)
	form.CustomDialog.Resize(formSize)
	(*form.parent).Resize(formSize)
	(*form.parent).Show()
	form.Show()
}

// Set the parent of the form.
func (form *Form) SetParent(parent *fyne.Window) {
	form.parent = parent
}

// Set the submit callback for the form.
func (form *Form) SetSubmitCallBack(callback Callback) {
	form.submitCallback = callback
}
