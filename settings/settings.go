package settings

import (
	"github.com/Polshkrev/gopolutils/fayl"
)

const (
	Folder   string      = "settings" // Parent folder of the application settings.
	Filename string      = "settings" // File name of the settings of the application.
	Suffix   fayl.Suffix = fayl.Toml  // Suffix of the settings of the application.
)

var (
	Path *fayl.Path = fayl.PathFromParts(Folder, Filename, Suffix) // Full path of the settings of the application.
)

// Settings for the header of the ui.
type Header struct {
	Width  uint8 `json:"width"`
	Height uint8 `json:"height"`
}

// Low-level configuration details.
type Configuration struct {
	Title         string `json:"title"`
	Header        Header `json:"header"`
	Width         uint16 `json:"width"`
	Height        uint16 `json:"height"`
	AppearanceKey string `json:"appearanceKey"`
}

// Settings of the application.
type Settings struct {
	Appearance    Appearance    `json:"appearance"`
	LocaleFolder  string        `json:"localeFolder"`
	Configuration Configuration `json:"configuration"`
}
