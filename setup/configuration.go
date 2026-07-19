package setup

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/settings"
)

// Setup the folders used by the application.
// If any of the given folders can not be setup, an [gopolutils.IOError] is returned.
func Configuration(configuration settings.Configuration) *gopolutils.Exception {
	var basePath *fayl.Path = GetBasePath(configuration)
	var except *gopolutils.Exception
	_, except = checkFile(basePath, createFolder)
	if except != nil {
		return except
	}
	var assetFolder *fayl.Path = basePath.JoinAs(configuration.AssetFolder)
	_, except = checkFile(assetFolder, createFolder)
	if except != nil {
		return except
	}
	var copyBlackPath *fayl.Path = assetFolder.JoinAs(configuration.BlackIcon)
	except = fayl.Write(copyBlackPath, copyImageDarkBytes)
	if except != nil {
		return except
	}
	var copyWhitePath *fayl.Path = assetFolder.JoinAs(configuration.WhiteIcon)
	except = fayl.Write(copyWhitePath, copyImageWhiteBytes)
	if except != nil {
		return except
	}
	var blackFile *fayl.Path = assetFolder.JoinAs(configuration.BlackLogo)
	except = fayl.Write(blackFile, blackLogoBytes)
	if except != nil {
		return except
	}
	var whiteFile *fayl.Path = assetFolder.JoinAs(configuration.WhiteLogo)
	except = fayl.Write(whiteFile, whiteLogoBytes)
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
