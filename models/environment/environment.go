package environment

import (
	"os"

	"github.com/Polshkrev/gopolutils"
)

// Check if the given key is within the system environment.
// Returns true if the given key is within the system environment, else false.
func Check(key string) bool {
	var result bool
	_, result = os.LookupEnv(key)
	return result
}

// Set a given key to a given value within the system environment.
// If the key can not be set in the system environment, an [gopolutils.OSError] is returned and the enviornment is not modified.
func Set(key, value string) *gopolutils.Exception {
	if Check(key) {
		return nil
	}
	var setError error = os.Setenv(key, value)
	if setError != nil {
		return gopolutils.NewNamedException(gopolutils.OSError, "%s", setError.Error())
	}
	return nil
}

// Obtain the value stored at the given key within the system environment.
// Returns the value stored at the given key within the system environment.
// If the key does not exist within the system environment, a [gopolutils.KeyError] is returned with an empty string.
func Get(key string) (string, *gopolutils.Exception) {
	if !Check(key) {
		return "", gopolutils.NewNamedException(gopolutils.KeyError, "Can not find \"%s\" in environment.", key)
	}
	return os.Getenv(key), nil
}
