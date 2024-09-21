package migration

import "os"

type MigrationFile interface {
	Create(database string, path string)
	GetMigrationSQLFiles() []os.DirEntry
	GetMigrationFileContent(file os.DirEntry) (string, error)
	DBMigrationFolderExists(folder string) bool
}
