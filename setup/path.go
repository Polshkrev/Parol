package setup

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

// Alias for a function to create a path.
type creationFunction[Type any] func(*fayl.Path) (*Type, *gopolutils.Exception)

// Check if the parent of the given path exists.
// Returns a tuple of the original path and a boolean on whether it exists.
func checkParent(path *fayl.Path) (*fayl.Path, bool) {
	var parent *fayl.Path = gopolutils.Must(path.Parent())
	return parent, parent.Exists()
}

// Create the parent of the given path if it does not already exist.
// If the given file can not be created, an [gopolutils.IOError] is returned.
func makeParent(path *fayl.Path) *gopolutils.Exception {
	var parent *fayl.Path
	var exists bool
	parent, exists = checkParent(path)
	if exists {
		return nil
	}
	var entry *fayl.Entry = fayl.NewEntry(parent)
	entry.SetType(fayl.DirectoryType)
	return entry.Create()
}

// Create a given path and its parent if they do no exist using a given creation function.
// Returns a pointer to the type created by the creation function.
// If the given file can not be created, an [gopolutils.IOError] is returned.
func checkFile[Type any](path *fayl.Path, creationFunction creationFunction[Type]) (*Type, *gopolutils.Exception) {
	var except *gopolutils.Exception = makeParent(path)
	if except != nil {
		return nil, except
	}
	return creationFunction(path)
}

// Create a given path.
// Returns the created orignal path.
// If the given path can not be created, an [gopolutils.IOError] is returned.
func createPath(path *fayl.Path) (*fayl.Path, *gopolutils.Exception) {
	if path.Exists() {
		return path, nil
	}
	return path, fayl.NewEntry(path).Create()
}

// Create a given folder.
// Returns the path of the created orignal folder.
// If the given folder can not be created, an [gopolutils.IOError] is returned.
func createFolder(path *fayl.Path) (*fayl.Path, *gopolutils.Exception) {
	if path.Exists() {
		return path, nil
	}
	var entry *fayl.Entry = fayl.NewEntry(path)
	entry.SetType(fayl.DirectoryType)
	return path, entry.Create()
}
