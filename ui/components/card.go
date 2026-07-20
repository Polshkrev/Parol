package components

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/ui/components/themes"
)

const (
	applicationName string = "Parol"     // Name of the current application.
	authorFolder    string = "polshkrev" // Author of the current application.
	assetsFolder    string = "assets"    // Folder where the assets are stored.
	copyDarkIcon    string = "black.png" // Black icon file name.
	copyLightIcon   string = "white.png" // White icon file name.
)

// Custom card widget.
type Card struct {
	*widget.Card
	title      string
	text       string
	appearance settings.Appearance
}

// Construct a new custom card widget based on a given title and appearance.
// Returns a new card based on a given title and appearance.
func NewCard(title, text string, appearance settings.Appearance) *Card {
	var card *Card = new(Card)
	card.title = title
	card.text = text
	card.appearance = appearance
	card.Card = &widget.Card{}
	card.ExtendBaseWidget(card.Card)
	return card
}

// Obtain the key of the card.
// Returns the title of the card.
func (card *Card) Key() string {
	return card.title
}

// Obtain the password of the card.
// Returns the text of the card.
func (card *Card) Password() string {
	return card.text
}

// Represent the card as a string.
// Returns a string representation of a card.
func (card *Card) String() string {
	return fmt.Sprintf("%s: %s", card.title, card.text)
}

// Paint a card to a given parent.
func (card *Card) Paint(parent *fyne.Container) {
	addCardBase(card, getCopyPath(card.appearance))
	registerCardEvents(parent)
	parent.Add(card)
	parent.Refresh()
}

// Construct the entire card base and paint it to its given parent.
func addCardBase(card *Card, assetPath *fayl.Path) {
	var copy *Button = gopolutils.Must(addCopyIcon(fayl.PathFrom(assetPath.String()), makeDefaultCallback(card.text)))
	var button *Button = NewButton("", theme.CancelIcon(), removeCallback(card))
	var cardTitle *fyne.Container = container.NewBorder(nil, nil, button, copy)
	var box *fyne.Container = container.NewVBox(container.NewPadded(cardTitle), layout.NewSpacer(), addLabel(card.title), layout.NewSpacer(), layout.NewSpacer())
	card.SetContent(box)
	container.NewThemeOverride(card, themes.NewCard())

}

// Construct the copy icon based on its given path and callback.
// Returns a button widget based on a given path and callback.
// If the file can not be loaded, an [gopolutils.IOError] is returned with a nil data pointer.
func addCopyIcon(file *fayl.Path, callback Callback) (*Button, *gopolutils.Exception) {
	var icon fyne.Resource
	var except error
	icon, except = fyne.LoadResourceFromPath(file.String())
	if except != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, "%s", except.Error())
	}
	return NewButton("", icon, callback), nil
}

// Construct a label based on a given text.
// Returns a container containing the label.
func addLabel(text string) *fyne.Container {
	var label *widget.Label = widget.NewLabel(text)
	return container.NewCenter(label)
}

// Obtain the path to the copy icon based on a given appearance.
// Returns the path to the copy icon based on a given appearance.
func getCopyPath(appearance settings.Appearance) *fayl.Path {
	var configurationDirectory *fayl.Path = fayl.Configuration()
	var authorFolder *fayl.Path = configurationDirectory.JoinAs(authorFolder)
	var applicationFolder *fayl.Path = authorFolder.JoinAs(applicationName)
	var assetFolder *fayl.Path = applicationFolder.JoinAs(assetsFolder)
	if appearance == settings.Light {
		return assetFolder.JoinAs(copyDarkIcon)
	}
	return assetFolder.JoinAs(copyLightIcon)
}
