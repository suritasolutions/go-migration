package migration

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/suritasolutions/go-migration/util"
)

func NewMigrationFile(ctx context.Context) *migrationFile {
	return &migrationFile{ctx}
}

type migrationFile struct {
	ctx context.Context
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
			util.Print("yellow", "Migrations folder already exists!")
		}
		return
	}

	if m.ctx.Value("verbose").(bool) {
		util.Print("gray", "Creating migrations folder...")
	}

	os.Mkdir("migrations", 0755)

	util.Print("green", "Migrations folder created successfully!")
}

func (m *migrationFile) dbMigrationFolderExists(folder string) bool {
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
			util.Print("yellow", "Migration folder already exists!")
		}
		return
	}

	if m.ctx.Value("verbose").(bool) {
		util.Print("gray", "Creating migration folder...")
	}

	os.Mkdir("migrations/"+folder, 0755)

	util.Print("green", "Migration folder created successfully!")
}

func (m *migrationFile) createMigrationFile(folder string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + folder + "/" + strings.Join(splittedPath, "/") + ".sql")
	defer file.Close()
	if err != nil {
		util.Print("red", "Error creating migration file")
		if m.ctx.Value("verbose").(bool) {
			util.Print("gray", err.Error())
		}
		return
	}

	util.Print("green", "Migration file created successfully!")
}

func (m *migrationFile) createRollbackFile(folder string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + folder + "/" + strings.Join(splittedPath, "/") + "_rollback.sql")
	defer file.Close()
	if err != nil {
		util.Print("red", "Error creating rollback file")
		if m.ctx.Value("verbose").(bool) {
			util.Print("gray", err.Error())
		}
		return
	}

	util.Print("green", "Rollback file created successfully!")
}

func (m *migrationFile) GetMigrationSQLFiles() []os.DirEntry {
	files, err := os.ReadDir("./migrations/" + m.ctx.Value("folder").(string))
	if err != nil {
		if m.ctx.Value("verbose").(bool) {
			util.Print("gray", err.Error())
		}
		util.Print("red", "Error reading migrations directory.")
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
