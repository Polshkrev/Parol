package components

import (
	"fyne.io/fyne/v2"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/parol/settings"
)

// Register the event triggered when a given card is deleted from its given parent.
func registerCardDeleted(parent *fyne.Container) {
	events.Subscribe(settings.CardDeleted, func(data any) {
		var card *Card
		var ok bool
		card, ok = data.(*Card)
		if !ok {
			return
		}
		parent.Remove(card)
		parent.Refresh()
	})
}

// Default callback for when a given card is deleted from its given parent.
// Returns a callback triggered when a given card is deleted from its given parent.
func removeCallback(card *Card) Callback {
	return func() {
		events.Post(settings.CardDeleted, card)
	}
}

// Register the card events.
func registerCardEvents(parent *fyne.Container) {
	registerCardDeleted(parent)
}
