package ui

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/ui/components"
)

// Register the card added event.
func registerCardAdded(gui *GUI) {
	events.Subscribe(settings.CardAdded, func(data any) {
		var card *components.Card
		var ok bool
		card, ok = data.(*components.Card)
		if !ok {
			return
		}
		gui.cards.Append(card)
		var except *gopolutils.Exception = gui.parol.Insert(card.Key(), card.Password())
		if except != nil {
			panic(except)
		}
	})
}

// Register the password deleted event.
func registerPasswordDelete(table table.Table[password.Password]) {
	events.Subscribe(settings.CardDeleted, func(data any) {
		var result *components.Card
		var ok bool
		result, ok = data.(*components.Card)
		if !ok {
			return
		}
		var password *password.Password = password.New(result.Key(), result.Password())
		table.Remove(*password)
	})
}

// Register the card deleted event.
func registerCardDeleted(gui *GUI) {
	events.Subscribe(settings.CardDeleted, func(data any) {
		var card *components.Card
		var ok bool
		card, ok = data.(*components.Card)
		if !ok {
			return
		}
		var except *gopolutils.Exception = removeCard(gui, card)
		if except != nil {
			panic(except)
		}
		except = gui.parol.Remove(card.Key())
		if except != nil {
			panic(except)
		}
	})
}

// Remove a given card from a given GUI.
// If the given card can not be removed from the given gui, a [gopolutils.ValueError] is returned.
func removeCard(gui *GUI, card components.Component) *gopolutils.Exception {
	var i gopolutils.Size
	for i = range collections.Enumerate(gui.cards) {
		var item *components.Component = gopolutils.Must(gui.cards.At(i))
		if (*item) != card {
			continue
		}
		return gui.cards.Remove(i)
	}
	return nil
}

// Register each of the card events for the given gui.
func registerCardEvents(gui *GUI) {
	registerCardAdded(gui)
	registerCardDeleted(gui)
	registerPasswordDelete(gui.passwords)
}
