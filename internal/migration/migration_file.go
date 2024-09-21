package migration

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/suritasolutions/go-migration/util"
)

func NewMigrationFile(ctx context.Context, logger util.Logger) MigrationFile {
	return &migrationFile{ctx, logger}
}

type migrationFile struct {
	ctx    context.Context
	logger util.Logger
}

func (m *migrationFile) Create(database string, path string) {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	m.createMigrationsFolder()
	m.createDatabaseFolder(database)
	m.createMigrationFile(database, path, timestamp)
	m.createRollbackFile(database, path, timestamp)
}

func (m *migrationFile) createMigrationsFolder() {
	_, err := os.Stat("./migrations")

	if !os.IsNotExist(err) {
		if m.ctx.Value("verbose").(bool) {
			m.logger.Debug("Migrations folder already exists!")
		}
		return
	}

	if m.ctx.Value("verbose").(bool) {
		m.logger.Debug("Creating migrations folder...")
	}

	os.Mkdir("migrations", 0755)

	m.logger.Success("Migrations folder created successfully!")
}

func (m *migrationFile) DBMigrationFolderExists(folder string) bool {
	if _, err := os.Stat("./migrations/" + folder); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (m *migrationFile) createDatabaseFolder(folder string) {
	_, err := os.Stat("./migrations/" + folder)

	if !os.IsNotExist(err) {
		if m.ctx.Value("verbose").(bool) {
			m.logger.Debug("Migration folder already exists!")
		}
		return
	}

	if m.ctx.Value("verbose").(bool) {
		m.logger.Debug("Creating migration folder...")
	}

	os.Mkdir("migrations/"+folder, 0755)

	m.logger.Success("Migration folder created successfully!")
}

func (m *migrationFile) createMigrationFile(folder string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + folder + "/" + strings.Join(splittedPath, "/") + ".sql")
	defer file.Close()
	if err != nil {
		m.logger.Fatal("Error creating migration file")
		if m.ctx.Value("verbose").(bool) {
			m.logger.Debug(err.Error())
		}
		return
	}

	m.logger.Success("Migration file created successfully!")
}

func (m *migrationFile) createRollbackFile(folder string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + folder + "/" + strings.Join(splittedPath, "/") + "_rollback.sql")
	defer file.Close()
	if err != nil {
		m.logger.Fatal("Error creating rollback file")
		if m.ctx.Value("verbose").(bool) {
			m.logger.Debug(err.Error())
		}
		return
	}

	m.logger.Success("Rollback file created successfully!")
}

func (m *migrationFile) GetMigrationSQLFiles() []os.DirEntry {
	files, err := os.ReadDir("./migrations/" + m.ctx.Value("folder").(string))
	if err != nil {
		m.logger.Error("Error reading migrations directory.")
		if m.ctx.Value("verbose").(bool) {
			m.logger.Debug(err.Error())
		}
	}

	migrations := []os.DirEntry{}
	for _, migration := range files {
		if !strings.HasSuffix(migration.Name(), "_rollback.sql") {
			migrations = append(migrations, migration)
		}
	}

	return migrations
}

func (m *migrationFile) GetMigrationFileContent(file os.DirEntry) (string, error) {
	content, err := os.ReadFile("./migrations/" + m.ctx.Value("folder").(string) + "/" + file.Name())
	if err != nil {
		return "", err
	}

	return string(content), nil
}
