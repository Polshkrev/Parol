package ui

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/events"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/ui/components"
)

// Register the card added event.
func registerCardAdded(gui *GUI, card *components.Card) {
	events.Subscribe(events.CardAdded, func() {
		gui.cards.Append(card)
		var except *gopolutils.Exception = gui.parol.Insert(card.Key(), card.Password())
		if except != nil {
			panic(except)
		}
		gui.passwords.Insert(*password.New(card.Key(), card.Password()))
	})
}

// Register the password deleted event.
func registerPasswordDelete(table table.Table[password.Password], password password.Password) {
	events.Subscribe(events.CardDeleted, func() {
		table.Remove(password)
	})
}

// Register the card deleted event.
func registerCardDeleted(gui *GUI, card *components.Card) {
	events.Subscribe(events.CardDeleted, func() {
		if gui.cards.IsEmpty() {
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
func registerCardEvents(gui *GUI, card *components.Card) {
	registerCardAdded(gui, card)
	registerCardDeleted(gui, card)
	registerPasswordDelete(gui.passwords, *password.New(card.Key(), card.Password()))
}
