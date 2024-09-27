package util

import (
	"io/fs"
	"os"
)

func NewStandardFileSystem() FileSystem {
	return &StandardFileSystem{}
}

type StandardFileSystem struct {
}

func (s *StandardFileSystem) FolderExists(folder string) (bool, error) {
	_, err := os.Stat("./migrations")
	if err != nil {
		return false, err
	}
	return !os.IsNotExist(err), nil
}

func (s *StandardFileSystem) CreateFolder(folder string, perm fs.FileMode) (bool, error) {
	err := os.Mkdir("migrations", 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}
