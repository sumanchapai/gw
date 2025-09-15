package scripts

import (
	"path/filepath"
	"runtime"

	"github.com/sumanchapai/gw/cerrors"
	"github.com/sumanchapai/gw/utils"
)

func scriptsDirPath() (string, error) {
	// Get the directory of this source file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", cerrors.ErrCantDetectCurrentPath
	}
	return filepath.Dir(filename), nil
}

// Returns the absolute path of the script with the given name
func Path(name string) (string, error) {
	dir, err := scriptsDirPath()
	if err != nil {
		return "", err
	}
	fullname := filepath.Join(dir, name)
	if err := utils.FileExists(fullname); err != nil {
		return "", err
	}
	return fullname, nil
}
