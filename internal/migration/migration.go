package migration

import (
	"context"
	"io/fs"
	"strings"

	"github.com/suritasolutions/go-migration/internal/db"
	"github.com/suritasolutions/go-migration/util"
)

func NewMigration(
	ctx context.Context,
	db db.Database,
	migrationFile MigrationFile,
	logger util.Logger,
) Migration {
	return &migration{ctx, db, migrationFile, logger}
}

type migration struct {
	ctx           context.Context
	db            db.Database
	migrationFile MigrationFile
	logger        util.Logger
}

func (m *migration) Migrate() {
	if !m.migrationFile.DBMigrationFolderExists(m.ctx.Value("folder").(string)) {
		m.logger.Fatal("The migration folder does not exist.")
		return
	}

	databaseExists, err := m.databaseExists()
	if err != nil {
		m.logger.Debug(err.Error())
		return
	}

	if !databaseExists {
		m.createDatabase()
	}

	migrationTableExists, err := m.migrationTableExists()
	if err != nil {
		m.logger.Fatal("Error checking if migration table exists.")
		m.logger.Debug(err.Error())
		return
	}

	if !migrationTableExists {
		if err := m.createMigrationTable(); err != nil {
			m.logger.Fatal("Error creating migration table.")
			m.logger.Debug(err.Error())
			return
		}
	}

	migrations := m.migrationFile.GetMigrationSQLFiles()

	for _, migration := range migrations {
		m.runMigrationSQLFile(migration)
	}
}

func (m *migration) Rollback() {
	// TODO: Implement rollback
}

func (m *migration) migrationTableExists() (bool, error) {
	db, err := m.db.ConnectDB()
	if err != nil {
		m.logger.Error("Error connecting to the database.")
		return false, err
	}
	defer db.Close()

	_, err = db.Exec("SELECT 1 FROM migrations LIMIT 1")
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (m *migration) createMigrationTable() error {
	db, err := m.db.ConnectDB()
	if err != nil {
		m.logger.Error("Error connecting to the database.")
		return err
	}
	defer db.Close()

	_, err = db.Exec(m.db.GetCreateMigrationTableSQL())
	if err != nil {
		return err
	}

	return nil
}

func (m *migration) databaseExists() (bool, error) {
	db, err := m.db.Connect()
	if err != nil {
		m.logger.Error("Error connecting to the database.")
		return false, err
	}
	defer db.Close()

	var tablename string

	scanErr := db.QueryRow(m.db.GetDatabaseExistsSQL()).Scan(&tablename)
	if scanErr != nil {
		if scanErr.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (m *migration) createDatabase() {
	db, err := m.db.Connect()
	if err != nil {
		m.logger.Fatal("Error connecting to the database.")
		m.logger.Debug(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec(m.db.GetCreateDatabaseSQL())
	if err != nil {
		m.logger.Fatal("Error creating database.")
		m.logger.Debug(err.Error())
		return
	}

	m.logger.Success("Database " + m.ctx.Value("database").(string) + " created.")
}

func (m *migration) runMigrationSQLFile(file fs.DirEntry) error {
	db, err := m.db.ConnectDB()
	if err != nil {
		m.logger.Error("Error connecting to the database")
		return err
	}
	defer db.Close()

	fileContent, err := m.migrationFile.GetMigrationFileContent(file)
	if err != nil {
		m.logger.Error("Error reading migration file content")
		m.logger.Debug(err.Error())
		return err
	}

	_, err = db.Exec(fileContent)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			m.logger.Info("Migration " + file.Name() + " already executed on database " + m.ctx.Value("database").(string) + ". Skipping...")
			return nil
		}

		m.logger.Error("Error running migration")
		m.logger.Debug(err.Error())
		return err
	}

	if err := m.registerMigrationExecutedOnMigrationsTable(file.Name()); err != nil {
		return err
	}

	m.logger.Success("Migration " + file.Name() + " executed on database " + m.ctx.Value("database").(string) + " successfully.")

	return nil
}

func (m *migration) registerMigrationExecutedOnMigrationsTable(fileName string) error {
	db, err := m.db.ConnectDB()
	if err != nil {
		m.logger.Error("Error connecting to the database")
		m.logger.Debug(err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO migrations (name) VALUES ($1)", fileName)
	if err != nil {
		m.logger.Error("Error registering migration on migrations table")
		m.logger.Debug(err.Error())
		return err
	}

	m.logger.Debug("Migration file " + fileName + " registered on migrations table")

	return nil
}
