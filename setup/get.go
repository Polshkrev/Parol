package setup

import (
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/parol/settings"
)

// Obtain the base path of the application.
// Returns the base path of the application.
func GetBasePath() *fayl.Path {
	var configurationDirectory *fayl.Path = fayl.Configuration()
	var authorFolder *fayl.Path = configurationDirectory.JoinAs(authorFolder)
	return authorFolder.JoinAs(applicationName)
}

// Obtain the asset path of the application.
// Returns the asset path of the application.
func GetAssetPath() *fayl.Path {
	return GetBasePath().JoinAs(assetsFolder)
}

// Obtain the settings path of the application.
// Returns the settings path of the application.
func GetSettingsPath() *fayl.Path {
	return GetBasePath().Join(*settings.Path)
}

// Obtain the data path of the application.
// Returns the data path of the application.
func GetDataFolder() *fayl.Path {
	return GetBasePath().JoinAs(dataFolder)
}

// Obtain the database path of the application.
// Returns the database path of the application.
func GetDatabaseFile() *fayl.Path {
	return GetDataFolder().JoinAs(databaseFile)
}

// Obtain the key path of the application.
// Returns the key path of the application.
func GetKeyFile() *fayl.Path {
	return GetDataFolder().JoinAs(keyFile)
}
