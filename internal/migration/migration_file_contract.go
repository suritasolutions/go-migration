package migration

import (
	"io/fs"
)

type MigrationFile interface {
	Create(database string, path string)
	GetMigrationSQLFiles() []fs.DirEntry
	GetMigrationFileContent(file fs.DirEntry) (string, error)
	DBMigrationFolderExists(folder string) bool
}
