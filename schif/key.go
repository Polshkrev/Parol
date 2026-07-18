package schif

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/fernet/fernet-go"
)

// Generate a key to encrypt data.
// Returns a fernetic key.
func GenerateKey() *fernet.Key {
	var key *fernet.Key = new(fernet.Key)
	var generateError error = key.Generate()
	if generateError != nil {
		panic(gopolutils.NewNamedException(gopolutils.KeyError, "%s", generateError.Error()))
	}
	return key
}

// Write a given key to a given [fayl.Path].
func WriteKey(file *fayl.Path, key *fernet.Key) {
	var except *gopolutils.Exception = fayl.Write(file, []byte(key.Encode()))
	if except != nil {
		panic(except)
	}
}

// Load a fernetic key from a given [fayl.Path].
// Returns a string representation of a fernetic key.
func LoadKey(file *fayl.Path) string {
	return string(gopolutils.Must(fayl.Read(file)))
}
