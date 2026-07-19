package setup

import (
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/settings"
)

// Obtain the base path of the application.
// Returns the base path of the application.
func GetBasePath(configuration settings.Configuration) *fayl.Path {
	var configurationDirectory *fayl.Path = fayl.Configuration()
	var authorFolder *fayl.Path = configurationDirectory.JoinAs(configuration.Author)
	return authorFolder.JoinAs(configuration.Title)
}

// Obtain the asset path of the application.
// Returns the asset path of the application.
func GetAssetPath(configuration settings.Configuration) *fayl.Path {
	return GetBasePath(configuration).JoinAs(configuration.AssetFolder)
}

// Obtain the settings path of the application.
// Returns the settings path of the application.
func GetSettingsPath(configuration settings.Configuration) *fayl.Path {
	return GetBasePath(configuration).Join(*settings.Path)
}

// Obtain the data path of the application.
// Returns the data path of the application.
func GetDataFolder(configuration settings.Configuration) *fayl.Path {
	return GetBasePath(configuration).JoinAs(configuration.DataFolder)
}

// Obtain the database path of the application.
// Returns the database path of the application.
func GetDatabaseFile(configuration settings.Configuration) *fayl.Path {
	return GetDataFolder(configuration).JoinAs(configuration.DatabaseFile)
}

// Obtain the key path of the application.
// Returns the key path of the application.
func GetKeyFile(configuration settings.Configuration) *fayl.Path {
	return GetDataFolder(configuration).JoinAs(configuration.KeyFile)
}
