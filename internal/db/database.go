package db

import "database/sql"

type Database interface {
	Connect() (*sql.DB, error)
	ConnectDB() (*sql.DB, error)
	GetCreateMigrationTableSQL() string
	GetDatabaseExistsSQL() string
	GetCreateDatabaseSQL() string
}
