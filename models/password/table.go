package password

import (
	"database/sql"
	"fmt"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/table"
)

var _ table.Table[Password] = (*Table)(nil)

const (
	TableName string = "passwords" // Name of the password table.
)

var (
	fieldString string = table.GetFields(key, value) // Representation of the column names within the table.
)

// Representation of a table of passwords.
type Table struct {
	connection *sql.DB
	name       string
}

// Construct a new table of passwords.
// Returns a new table of passwords.
func NewTable(connection *sql.DB) *Table {
	var table *Table = new(Table)
	table.SetConnection(connection)
	return table
}

// Set the name of the table.
func (storage *Table) SetName(name string) {
	storage.name = name
}

// Set the connection to the table.
func (storage *Table) SetConnection(connection *sql.DB) {
	storage.connection = connection
}

// Obtain the name of the table.
// Returns the table name.
func (storage Table) Name() string {
	return storage.name
}

// Obtain the connection to the table.
// Returns the connection to the table.
func (storage Table) Connection() *sql.DB {
	return storage.connection
}

// Create a table with a given name.
// If the table can not be created, an [gopolutils.IOError] is returned.
func (storage *Table) Create(name string) *gopolutils.Exception {
	var execError error
	_, execError = storage.connection.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s TEXT NOT NULL, %s TEXT NOT NULL, UNIQUE(%s));", name, key, value, key))
	if execError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", execError.Error())
	}
	storage.SetName(name)
	return nil
}

// Insert a given encrypted password into the table.
// If the password can not be inserted, an [gopolutils.IOError] is returned.
func (storage *Table) Insert(password Password) *gopolutils.Exception {
	var execError error
	_, execError = storage.connection.Exec(fmt.Sprintf("INSERT OR IGNORE INTO %s (%s) VALUES (?, ?);", storage.Name(), fieldString), password.Key(), password.Password())
	if execError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", execError.Error())
	}
	return nil
}

// Insert multiple encrypted passwords into the table.
// If any of the passwords can not be inserted, an [gopolutils.IOError] is returned.
func (storage *Table) InsertMany(passwords collections.View[Password]) *gopolutils.Exception {
	var i int
	for i = range passwords.Collect() {
		var item Password = passwords.Collect()[i]
		var execError *gopolutils.Exception = storage.Insert(item)
		if execError != nil {
			return execError
		}
	}
	return nil
}

// Obtain a password with a given id from the table.
// Returns a password stored in the table at the given id.
//
// This method is not implemented yet and will always return a [gopolutils.NotImplementedError] with a nil data pointer.
func (storage *Table) Get(id gopolutils.Size) (*Password, *gopolutils.Exception) {
	return nil, gopolutils.NewNamedException(gopolutils.NotImplementedError, "Get by id has not been implemented yet.")
}

// Append a given table row to a given collection of passwords.
// If the rows can not be scanned, an [gopolutils.IOError] is returned.
func appendPassword(passwords collections.Collection[*Password], rows *sql.Rows) *gopolutils.Exception {
	var key, password string
	var scanError error = rows.Scan(&key, &password)
	if scanError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", scanError.Error())
	}
	var result *Password = New(key, password)
	passwords.Append(result)
	return nil
}

// Obtain all the passwords stored in the table.
// Returns a [collections.View] of [Password]s strored in the table.
// If the passwords can not be obtained, an [gopolutils.IOError] is returned with a nil data pointer.
func (storage *Table) GetAll() (collections.View[*Password], *gopolutils.Exception) {
	var result collections.Collection[*Password] = collections.NewArray[*Password]()
	var rows *sql.Rows
	var queryError error
	rows, queryError = storage.connection.Query(fmt.Sprintf("SELECT * FROM %s", storage.Name()))
	if queryError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, "%s", queryError.Error())
	}
	defer rows.Close()
	for rows.Next() {
		appendPassword(result, rows)
	}
	return result, nil
}

// Remove a given password from the table.
// If the given password can not be removed from the table, an [gopolutils.IOError] is returned.
func (storage *Table) Remove(item Password) *gopolutils.Exception {
	var removeError error
	_, removeError = storage.connection.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = ?;", storage.Name(), key), item.Key())
	if removeError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", removeError.Error())
	}
	return nil
}

// Drop a table with a given name.
// If the table can not be dropped, an [gopolutils.IOError] is returned.
func (storage *Table) Drop(name string) *gopolutils.Exception {
	var dropError error
	_, dropError = storage.connection.Exec(fmt.Sprintf("DROP TABLE %s;", name))
	if dropError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", dropError.Error())
	}
	return nil
}

// Close the connection to the table.
// If the connection to the table can not be closed, an [gopolutils.IOError] is returned.
func (storage *Table) Close() *gopolutils.Exception {
	var closeError error = storage.connection.Close()
	if closeError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, "%s", closeError.Error())
	}
	return nil
}

// Obtain the size of the table.
// Returns the size of the table.
func (storage *Table) Size() gopolutils.Size {
	var count gopolutils.Size
	var row *sql.Row = storage.connection.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s;", storage.Name()))
	var countError error = row.Scan(&count)
	if countError != nil {
		panic(gopolutils.NewNamedException(gopolutils.IOError, "%s", countError.Error()))
	}
	return count
}

// Determine if the table is empty.
// Returns true if the size of the table is equal to zero, else false.
func (storage *Table) IsEmpty() bool {
	return storage.Size() == 0
}
