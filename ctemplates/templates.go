package ctemplates

import (
	"path/filepath"
	"runtime"

	"github.com/sumanchapai/gw/cerrors"
)

func templateDirPath() (string, error) {
	// Get the directory of this source file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", cerrors.ErrCantDetectCurrentPath
	}
	return filepath.Dir(filename), nil
}

// Returns the absolute path of the template with the given name
func Path(name string) (string, error) {
	dir, err := templateDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, name), nil
}
