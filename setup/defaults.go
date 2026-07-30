package setup

import (
	"github.com/Polshkrev/goserialize"
)

const (
	applicationName string = "Parol"          // Name of the current application.
	authorFolder    string = "polshkrev"      // Author of the current application.
	assetsFolder    string = "assets"         // Folder where the assets are stored.
	copyBlackIcon   string = "black.png"      // Black icon file name.
	copyWhiteIcon   string = "white.png"      // White icon file name.
	blackLogo       string = "black_logo.png" // Black logo file.
	whiteLogo       string = "white_logo.png" // White Logo file.
	dataFolder      string = "data"           // Data folder.
	databaseFile    string = "database.db"    // Database filename.
	keyFile         string = "Key.key"        // Key filename.
)

// Default settings for the application.
var defaultSettings goserialize.Object = goserialize.Object{
	"appearance": "system",
	"configuration": goserialize.Object{
		"title": "Parol",
		"header": goserialize.Object{
			"width":  50,
			"height": 50,
		},
		"appearanceKey": "FYNE_THEME",
		"width":         1500,
		"height":        750,
	},
}
