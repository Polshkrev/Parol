package components

import (
	"fyne.io/fyne/v2"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/parol/settings"
)

// Register the event triggered when a given card is deleted from its given parent.
func registerCardDeleted(card *Card, parent *fyne.Container) {
	events.Subscribe(settings.CardDeleted, func() {
		removeCallback(card, parent)
	})
}

// Default callback for when a given card is deleted from its given parent.
// Returns a callback triggered when a given card is deleted from its given parent.
func removeCallback(card *Card, parent *fyne.Container) Callback {
	return func() {
		parent.Remove(card)
		parent.Refresh()
		events.Post(settings.CardDeleted)
	}
}

// Register the card events.
func registerCardEvents(card *Card, parent *fyne.Container) {
	registerCardDeleted(card, parent)
}
