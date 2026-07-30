package main

import (
	"flag"
	"fmt"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/models/environment"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/schif/parol"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/setup"
	"github.com/Polshkrev/parol/ui"
	"github.com/thiagokokada/dark-mode-go"
)

const (
	driver table.Driver = table.Sqlite // Default database driver.
)

// Register the application end event.
func registerApplicationEnd(manager *parol.Parol, database table.Table[password.Password]) {
	events.Subscribe(settings.ApplicationEnd, func(any) {
		var except *gopolutils.Exception = database.InsertMany(parol.ObjectToView(manager.Passwords()))
		if except != nil {
			panic(except)
		}
	})
}

// Obtain the system appearance.
// Returns the [settings.Appearance] of the system.
func systemMode() settings.Appearance {
	var isDark bool
	var darkError error
	isDark, darkError = dark.IsDarkMode()
	if darkError != nil {
		panic(gopolutils.NewNamedException(gopolutils.OSError, "%s", darkError.Error()))
	} else if isDark {
		return settings.Dark
	}
	return settings.Light
}

// Obtain the system appearance if the given appearance is not already set.
// Returns the system appearance if the given appearance is not already set.
func systemModeFrom(appearance settings.Appearance) settings.Appearance {
	if appearance != settings.System {
		return appearance
	}
	return systemMode()
}

// Set the theme environment variable based on the given settings.
func setTheme(configuration *settings.Settings) {
	var appearance settings.Appearance = systemModeFrom(configuration.Appearance)
	var except *gopolutils.Exception = environment.Set(configuration.Configuration.AppearanceKey, appearance.String())
	if except != nil {
		panic(except)
	}
	configuration.Appearance = appearance
}

// Obtain a [collections.View] of [password.Password]s based on the given variadic keys.
// Returns a [collections.View] of [password.Password]s based on the given variadic keys.
func passwordsFrom(parol parol.Parol, arguments ...string) collections.View[password.Password] {
	var result collections.Collection[password.Password] = collections.NewArray[password.Password]()
	var i int
	for i = range arguments {
		var key string = arguments[i]
		var value []byte = gopolutils.Must(parol.Get(key))
		var password *password.Password = password.New(key, string(value))
		result.Append(*password)
	}
	return result
}

// Print each of the given [password.Password]s.
func printPasswords(passwords collections.View[password.Password]) {
	var i int
	for i = range passwords.Collect() {
		fmt.Println(passwords.Collect()[i])
	}
}

func main() {
	flag.Parse()
	var manager *parol.Parol
	var passwords table.Table[password.Password]
	manager, passwords = setup.Passwords(setup.GetKeyFile(), setup.GetDatabaseFile(), password.TableName, driver)
	defer passwords.Close()
	var configuration *settings.Settings = gopolutils.Must(fayl.ReadObject[settings.Settings](setup.GetSettingsPath()))
	setTheme(configuration)
	registerApplicationEnd(manager, passwords)
	if len(flag.Args()) == 0 {
		var gui *ui.GUI = setup.GUI(passwords, manager, *configuration)
		gui.Run()
	} else if len(flag.Args()) == 2 {
		var key string = flag.Arg(0)
		var password string = flag.Arg(1)
		var except *gopolutils.Exception = manager.Insert(key, password)
		if (except != nil) && (!except.Is(gopolutils.KeyError)) {
			panic(except)
		}
		except = passwords.InsertMany(parol.ObjectToView(manager.Passwords()))
		if except != nil {
			panic(except)
		}
	} else {
		var values collections.View[password.Password] = passwordsFrom(*manager, flag.Args()...)
		printPasswords(values)
	}
}
