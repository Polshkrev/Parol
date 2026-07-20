package setup

import (
	"database/sql"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/gopolutils/table/connect"
	"github.com/Polshkrev/parol/models/password"
	"github.com/Polshkrev/parol/schif/key"
	"github.com/Polshkrev/parol/schif/parol"
	"github.com/fernet/fernet-go"
	_ "github.com/mattn/go-sqlite3"
)

// Decode the key loaded from a given [fayl.Path].
// Returns a [fernet.Key] from the given [fayl.Path].
// If the key can not be loaded, a [gopolutils.IOError] is returned with a nil data pointer.
func loadKey(file *fayl.Path) (*fernet.Key, *gopolutils.Exception) {
	var result *fernet.Key
	var keyError error
	result, keyError = fernet.DecodeKey(key.Load(file))
	if keyError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.KeyError, "%s", keyError.Error())
	}
	return result, nil
}

// Generate and write a [fernet.Key] to a given [fayl.Path].
// Returns a [fernet.Key].
// If the key can not be loaded or written, a [gopolutils.IOError] is returned with a nil data pointer.
func keygen(file *fayl.Path) (*fernet.Key, *gopolutils.Exception) {
	if file.Exists() {
		return loadKey(file)
	}
	var result *fernet.Key = key.Generate()
	key.Write(file, result)
	return result, nil
}

// Setup the password logic for the application.
// Returns a tuple containing the password manager and the table of passwords.
func Passwords(keyFile, passwordFile *fayl.Path, tableName string, driver table.Driver) (*parol.Parol, table.Table[password.Password]) {
	var except *gopolutils.Exception = Configuration(applicationName)
	if except != nil {
		panic(except)
	}
	var key *fernet.Key = gopolutils.Must(checkFile(keyFile, keygen))
	gopolutils.Must(checkFile(passwordFile, createPath))

	var parol *parol.Parol = parol.New(key)

	var connection *sql.DB = gopolutils.Must(connect.Connect(driver, passwordFile))

	var database table.Table[password.Password] = password.NewTable(connection)
	database.Create(tableName)

	except = parol.Extend(gopolutils.Must(database.GetAll()))
	if except != nil {
		panic(except)
	}

	return parol, database
}
