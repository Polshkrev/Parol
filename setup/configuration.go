package setup

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/ui"
)

// Setup the folders used by the application.
// If any of the given folders can not be setup, an [gopolutils.IOError] is returned.
func Configuration(name string) *gopolutils.Exception {
	var basePath *fayl.Path = GetBasePath()
	var except *gopolutils.Exception
	_, except = checkFile(basePath, createFolder)
	if except != nil {
		return except
	}
	var assetFolder *fayl.Path = basePath.JoinAs(assetsFolder)
	_, except = checkFile(assetFolder, createFolder)
	if except != nil {
		return except
	}
	var copyBlackPath *fayl.Path = assetFolder.JoinAs(copyBlackIcon)
	except = fayl.Write(copyBlackPath, ui.BlackCopy)
	if except != nil {
		return except
	}
	var copyWhitePath *fayl.Path = assetFolder.JoinAs(copyWhiteIcon)
	except = fayl.Write(copyWhitePath, ui.WhiteCopy)
	if except != nil {
		return except
	}
	var blackFile *fayl.Path = assetFolder.JoinAs(blackLogo)
	except = fayl.Write(blackFile, ui.BlackLogo)
	if except != nil {
		return except
	}
	var whiteFile *fayl.Path = assetFolder.JoinAs(whiteLogo)
	except = fayl.Write(whiteFile, ui.WhiteLogo)
	if except != nil {
		return except
	}
	var settingsPath *fayl.Path = basePath.Join(*settings.Path)
	if settingsPath.Exists() {
		return nil
	}
	gopolutils.Must(checkFile(settingsPath, createPath))
	return fayl.WriteObject(settingsPath, &defaultSettings)
}
