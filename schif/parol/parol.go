package parol

import (
	"cmp"
	"slices"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/goserialize"
	"github.com/Polshkrev/parol/models/password"
	"github.com/fernet/fernet-go"
)

// A password manager.
type Parol struct {
	passwords collections.Mapping[string, []byte]
	key       *fernet.Key
}

// Construct a new password manager based on a given key.
// Returns a new password manager based on a given key.
func New(key *fernet.Key) *Parol {
	var parol *Parol = new(Parol)
	parol.passwords = collections.NewMap[string, []byte]()
	parol.key = key
	return parol
}

// Insert a given password into the manager by its given key.
// If the key is already in the manager, instead of just quietly not inserting into the manager, a [gopolutils.KeyError] is returned.
// If a [gopolutils.KeyError] is returned, the manager is not modified.
func (parol *Parol) Insert(key, password string) *gopolutils.Exception {
	var token []byte = gopolutils.Must(encryptPassword(parol.key, password))
	return parol.passwords.Insert(key, token)
}

// Append a [collections.View] of [password.Password]s to the manager.
// If the password can not be appended to the manager, a [gopolutils.KeyError] is returned.
func (parol *Parol) Extend(passwords collections.View[*password.Password]) *gopolutils.Exception {
	var raw []*password.Password = passwords.Collect()
	var i int
	for i = range raw {
		var password *password.Password = raw[i]
		var except *gopolutils.Exception = parol.passwords.Insert(password.Key(), []byte(password.Password()))
		if except != nil {
			return except
		}
	}
	return nil
}

// Obtain the password stored in the manager at the given key.
// Returns the decrypted password stored in the manager at the given key.
// If the manager is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If the key is not in the manager, a [gopolutils.KeyError] is returned with a nil data pointer.
// If the given password can not be decrypted, a [gopolutils.KeyError] is returned with a nil data pointer.
func (parol Parol) Get(key string) ([]byte, *gopolutils.Exception) {
	var token *[]byte
	var except *gopolutils.Exception
	token, except = parol.passwords.At(key)
	if except != nil {
		return nil, except
	}
	return decryptPassword(*token, parol.key)
}

// Collect the passwords stored in the manager.
// Returns a [goserialize.Object] of each of the key-value pair stored in the manager.
func (parol Parol) Passwords() goserialize.Object {
	var bucket collections.Pair[string, []byte]
	var result goserialize.Object = make(goserialize.Object, 0)
	for _, bucket = range parol.passwords.Collect() {
		var key *string
		var value *[]byte
		key, value = bucket.Items()
		result[*key] = string(*value)
	}
	return result
}

// Remove a password in the manager at the given key.
// If the manager is empty, a [gopolutils.ValueError] is returned.
// If the given key is not stored in the manager, a [gopolutils.KeyError] is returned.
// If a [gopolutils.ValueError] or a [gopolutils.KeyError] is returned, the manager is not modified.
func (parol *Parol) Remove(key string) *gopolutils.Exception {
	return parol.passwords.Remove(key)
}

// Determine if the given key exists within the mapping.
// Returns true if the mapping contains the given key.
func (parol Parol) HasKey(key string) bool {
	return parol.passwords.HasKey(key)
}

// Encrypt a given password using a given key.
// Returns a slice of bytes containing the encrypted password.
// If the password can not be encrypted, a [gopolutils.ValueError] is returned with a nil data pointer.
func encryptPassword(key *fernet.Key, password string) ([]byte, *gopolutils.Exception) {
	var token []byte
	var encryptError error
	token, encryptError = fernet.EncryptAndSign([]byte(password), key)
	if encryptError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "%s", encryptError.Error())
	}
	return token, nil
}

// Decrypt a given password using a given key.
// Returns a slice of bytes containing the decrypted password.
// If the given password can not be decrypted, a [gopolutils.KeyError] is returned with a nil data pointer.
func decryptPassword(password []byte, key *fernet.Key) ([]byte, *gopolutils.Exception) {
	var result []byte = fernet.VerifyAndDecrypt(password, 0, []*fernet.Key{key})
	if result == nil {
		return nil, gopolutils.NewNamedException(gopolutils.KeyError, "Can not decrypt password \"%s\" at key \"%s\"", string(password), key)
	}
	return result, nil
}

// Sort a given [collections.View] of [password.Password]s.
// Returns a sorted copy of the given [collections.View] of [password.Password]s.
func sort(values collections.Collection[password.Password]) collections.View[password.Password] {
	var raw []password.Password = values.Collect()
	var result collections.Collection[password.Password] = collections.NewArray[password.Password]()
	slices.SortFunc(raw, func(a, b password.Password) int {
		return cmp.Compare(a.Key(), b.Key())
	})
	var i int
	for i = range raw {
		result.Append(raw[i])
	}
	return result
}

// Convert a given [goserialize.Object] into a [collections.View] of [password.Password]s.
// Returns [collections.View] of [password.Password]s containing each of the values of the given [goserialize.Object].
func ObjectToView(object goserialize.Object) collections.View[password.Password] {
	var result collections.Collection[password.Password] = collections.NewArray[password.Password]()
	var key string
	var value any
	for key, value = range object {
		result.Append(*password.New(key, value.(string)))
	}
	return sort(result)
}
