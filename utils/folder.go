package utils

import (
	"os"

	"github.com/sumanchapai/gw/cerrors"
)

// Returns nil if the folder exists, err otherwise
func FolderExists(name string) error {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return cerrors.ErrPathDoesntExist
	}
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return cerrors.ErrPathNotAFolder
	}
	return nil
}
