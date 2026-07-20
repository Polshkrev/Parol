package setup

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/schif/parol"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/ui"
)

// Load the icon based on the given appearance.
// Returns a [fyne.Resource] of the icon based on the given appearance.
func loadIcon(appearance settings.Appearance, whitePath, darkPath *fayl.Path) fyne.Resource {
	var resource fyne.Resource
	var resourceError error
	if appearance == settings.Dark {
		resource, resourceError = fyne.LoadResourceFromPath(darkPath.String())
		if resourceError != nil {
			panic(gopolutils.NewNamedException(gopolutils.IOError, "%s", resourceError.Error()))
		}
	} else {
		resource, resourceError = fyne.LoadResourceFromPath(whitePath.String())
		if resourceError != nil {
			panic(gopolutils.NewNamedException(gopolutils.IOError, "%s", resourceError.Error()))
		}
	}
	return resource
}

// Setup the gui based on its parts.
// Returns the gui of the application.
func GUI(passwords table.Table[password.Password], manager *parol.Parol, configuration settings.Settings) *ui.GUI {
	var application fyne.App = app.New()
	application.SetIcon(loadIcon(configuration.Appearance, GetAssetPath().JoinAs(whiteLogo), GetAssetPath().JoinAs(blackLogo)))
	var window fyne.Window = application.NewWindow(configuration.Configuration.Title)
	var gui *ui.GUI = ui.NewGUI(&window, passwords, manager, configuration)
	gui.Paint()
	return gui
}
