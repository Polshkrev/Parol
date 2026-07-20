package password

import (
	"fmt"

	"github.com/Polshkrev/gopolutils/table"
)

const (
	key   table.Field = "key"      // Key column name.
	value table.Field = "password" // Password column name.
)

// Representation of a password with a key and value.
type Password struct {
	key      string
	password string
}

// Construct a new password based on its key and value.
// Returns a new password of the given key and value.
func New(key, password string) *Password {
	var result *Password = new(Password)
	result.key = key
	result.password = password
	return result
}

// Set the key of the password.
func (password *Password) SetKey(key string) {
	password.key = key
}

// Set the password of the password.
func (password *Password) SetPassword(value string) {
	password.password = value
}

// Obtain the key of the password.
// Returns the key of the password.
func (password Password) Key() string {
	return password.key
}

// Obtain the value of the password.
// Returns the value of the password.
func (password Password) Password() string {
	return password.password
}

// Represent the password as a string.
// Returns a string representation of a password.
func (password Password) String() string {
	return fmt.Sprintf("%s - %s", password.key, password.password)
}
