package ui

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/models/environment"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/schif/parol"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/ui/components"
	"github.com/kbinani/screenshot"
)

// Custom GUI.
type GUI struct {
	settings  settings.Settings
	parol     *parol.Parol
	window    *fyne.Window
	passwords table.Table[password.Password]
	cards     collections.Collection[components.Component]
}

// Construct a gui based on its parts.
// Returns the application gui.
func NewGUI(window *fyne.Window, passwords table.Table[password.Password], parol *parol.Parol, settings settings.Settings) *GUI {
	var gui *GUI = new(GUI)
	gui.window = window
	gui.parol = parol
	gui.settings = settings
	gui.passwords = passwords
	gui.cards = collections.NewArray[components.Component]()
	return gui
}

// Paint the gui.
func (gui *GUI) Paint() {
	var windowSize fyne.Size = getWindowSize(gui.settings.Configuration)
	(*gui.window).SetContent(paint(gui))
	(*gui.window).Resize(windowSize)
	(*gui.window).CenterOnScreen()
	events.Subscribe(settings.ApplicationEnd, func(any) {
		(*gui.window).Close()
	})
	(*gui.window).SetCloseIntercept(func() {
		events.Post(settings.ApplicationEnd, nil)
	})
}

// Show and run the gui.
func (gui *GUI) Run() {
	(*gui.window).ShowAndRun()
}

// Setup the gui based on its given parent container.
func (gui *GUI) setup(parent *fyne.Container) {
	var key string
	for key = range gui.parol.Passwords() {
		var card *components.Card = components.NewCard(key, string(gopolutils.Must(gui.parol.Get(key))), gui.settings.Appearance)
		card.Paint(parent)
		gui.cards.Append(card)
	}
	registerCardEvents(gui)
}

// Setup the base ui of the given gui.
// Returns a [fyne.Container] based on each of its components.
func paint(gui *GUI) *fyne.Container {
	var gridSize fyne.Size = fyne.NewSize(float32((gui.settings.Configuration.Width/5)-5), float32((gui.settings.Configuration.Height/5)+100))
	var content *fyne.Container = container.New(layout.NewGridWrapLayout(gridSize))
	gui.setup(content)
	var button *components.Button = components.NewButton("", theme.ContentAddIcon(), func() {
		paintAddForm(gui, content, "Add Password", gui.settings.Appearance)
	})
	var label *widget.Label = widget.NewLabelWithStyle("Passwords:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
	var header *components.Header = components.NewHeader(label, button, gui.settings.Configuration.Header.Width, gui.settings.Configuration.Header.Height)
	header.Paint(content)
	return container.NewBorder(header.Parent(), nil, nil, nil, container.NewVScroll(content))
}

// Paint the add form based on its parent title, and appearance.
func paintAddForm(gui *GUI, cardParent *fyne.Container, title string, appearance settings.Appearance) {
	var formWindow fyne.Window = fyne.CurrentApp().NewWindow(title)
	var form *components.Form
	form = components.NewForm("Add Password", "Add", "Cancel", 300, 300, nil, onFormCancel(formWindow))
	form.SetParent(&formWindow)
	var callback components.Callback = registerAddFormSubmit(gui, cardParent, form, appearance)
	form.Append(components.NewItem(components.NewEntry("Key", false, false, func(string) error { return nil }, callback)))
	form.Append(components.NewItem(components.NewEntry("Password", true, true, func(string) error { return nil }, callback)))
	form.SetSubmitCallBack(callback)
	form.Paint(nil)
}

// Paint the information dialog based on a title and message.
func paintError(title, message string, width, height uint16) {
	var parent fyne.Window = fyne.CurrentApp().NewWindow(title)
	var information dialog.Dialog = dialog.NewInformation(title, message, parent)
	information.SetOnClosed(func() {
		parent.Close()
	})
	var infoSize fyne.Size = fyne.NewSize(float32(width), float32(height))
	information.Resize(infoSize)
	parent.Resize(infoSize)
	information.Show()
	parent.CenterOnScreen()
	parent.RequestFocus()
	parent.SetFixedSize(true)
	parent.Show()
}

// Default callback for the form based on its given parent.
// Returns the default callback triggered when the form is canceled.
func onFormCancel(parent fyne.Window) components.Callback {
	return func() {
		parent.Close()
	}
}

// Obtain the form items.
// Returns a [collections.Pair] of each of the form items.
func getFormItems(form *components.Form) *collections.Pair[*components.Item, *components.Item] {
	return collections.NewPair(form.Items().Collect()[0], form.Items().Collect()[1])
}

// Setup the logic for submitting a form.
// Returns a callback triggered when the form is submitted.
func registerAddFormSubmit(gui *GUI, parent *fyne.Container, form *components.Form, appearance settings.Appearance) components.Callback {
	return func() {
		var items *collections.Pair[*components.Item, *components.Item] = getFormItems(form)
		var card *components.Card = components.NewCard((*items.First()).Value(), (*items.Second()).Value(), appearance)
		if len(card.Key()) == 0 {
			paintError("Empty Key", "The Key Entry Can Not Be Empty.", 400, 150)
			return
		} else if len(card.Password()) == 0 {
			paintError("Empty Password", "The Password Entry Can Not Be Empty.", 400, 150)
			return
		} else if gui.parol.HasKey(card.Key()) {
			paintError("Duplicate Key", "Can not add a duplicate key.", 400, 150)
			return
		}
		card.Paint(parent)
		events.Post(settings.CardAdded, card)
		(*form.Parent()).Close()
	}
}

// Scale the application based on the original size.
// Returns a size based on a scaled output.
func getWindowSize(configuration settings.Configuration) fyne.Size {
	var bounds image.Rectangle = screenshot.GetDisplayBounds(0)
	if bounds.Dx() > 1920 && bounds.Dx() > 1080 {
		environment.Set("FYNE_SCALE", ".825")
	}
	return fyne.NewSize(float32(configuration.Width), float32(configuration.Height))
}
