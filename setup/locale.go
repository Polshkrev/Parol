package setup

import (
	"fyne.io/fyne/v2/lang"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/locale"
)

// Setup the localization based on the given folder.
// If the localization can not be loaded, an [gopolutils.OSError] is returned.
func Locale(folder *fayl.Path) *gopolutils.Exception {
	var embedError error = lang.AddTranslationsFS(locale.LocalFS, folder.String())
	if embedError != nil {
		return gopolutils.NewNamedException(gopolutils.OSError, "%s", embedError.Error())
	}
	return nil
}
