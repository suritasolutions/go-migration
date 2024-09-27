package util

import "io/fs"

type FileSystem interface {
	FolderExists(folder string) (bool, error)
	CreateFolder(folder string, perm fs.FileMode) (bool, error)
}
