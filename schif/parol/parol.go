package schif

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/goserialize"
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
	return parol.passwords.Insert(key, gopolutils.Must(encryptPassword(parol.key, password)))
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
	return decryptPassword(parol.key, *token)
}

// Collect the passwords stored in the manager.
// Returns a [goserialize.Object] of each of the key-value pair stored in the manager.
func (parol Parol) Passwords() goserialize.Object {
	var result goserialize.Object = make(goserialize.Object, 0)
	var i gopolutils.Size
	for i = range collections.Enumerate(parol.passwords) {
		var bucket collections.Pair[string, []byte] = parol.passwords.Collect()[i]
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
func decryptPassword(key *fernet.Key, password []byte) ([]byte, *gopolutils.Exception) {
	var result []byte = fernet.VerifyAndDecrypt(password, 0, []*fernet.Key{key})
	if result == nil {
		return nil, gopolutils.NewNamedException(gopolutils.KeyError, "Can not decrypt password \"%s\" at key \"%s\"", string(password), key)
	}
	return result, nil
}
