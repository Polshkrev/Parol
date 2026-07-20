package settings

import "github.com/Polshkrev/gopolutils/events"

const (
	ApplicationStart events.EventType = "applicationStart" // Name of the application start event type.
	CardAdded        events.EventType = "cardAdded"        // Name of the card added event type.
	CardDeleted      events.EventType = "cardDeleted"      // Name of the card deleted event type.
	ApplicationEnd   events.EventType = "applicationEnd"   // Name of the application end event type.
)
