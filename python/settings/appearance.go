package settings

import (
	"github.com/Polshkrev/gopolutils"
)

// Appearance of the ui of the application.
type Appearance gopolutils.StringEnum

const (
	Light  Appearance = "light"  // Light appearance of the ui.
	Dark   Appearance = "dark"   // Dark appearance of the ui.
	System Appearance = "system" // Appearance of the system.
)

// Represent the appearance as a string.
// Returns a string representation of the appearance.
func (appearance Appearance) String() string {
	return string(appearance)
}
