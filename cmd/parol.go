package main

import (
	"flag"
	"fmt"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/events"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/parol/models/environment"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/schif/parol"
	"github.com/Polshkrev/parol/settings"
	"github.com/Polshkrev/parol/setup"
	"github.com/Polshkrev/parol/ui"
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

// Set the theme environment variable based on the given settings.
func setTheme(setttings settings.Settings) {
	var except *gopolutils.Exception = environment.Set(setttings.Configuration.AppearanceKey, setttings.Appearance.String())
	if except != nil {
		panic(except)
	}
}

func main() {
	var manager *parol.Parol
	var passwords table.Table[password.Password]
	manager, passwords = setup.Passwords(setup.GetKeyFile(), setup.GetDatabaseFile(), password.TableName, driver)
	defer passwords.Close()
	var configuration *settings.Settings = gopolutils.Must(fayl.ReadObject[settings.Settings](setup.GetSettingsPath()))
	setTheme(*configuration)
	registerApplicationEnd(manager, passwords)
	if len(flag.Args()) == 0 {
		var gui *ui.GUI = setup.GUI(passwords, manager, *configuration)
		gui.Run()
	} else if len(flag.Args()) < 2 {
		var password []byte = gopolutils.Must(manager.Get(flag.Arg(0)))
		fmt.Println(string(password))
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
	}
}
