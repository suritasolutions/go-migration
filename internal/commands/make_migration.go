package commands

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var MakeMigrationCmd = &cobra.Command{
	Use:   "make:migration",
	Short: "Generate a new migration",
	Long:  `Generate a new migration file.`,
	Run:   handle,
}

func init() {
	MakeMigrationCmd.SetUsageTemplate(`Usage:
    - migration make:migration [database] [migration_name]

    `)
	rootCmd.AddCommand(MakeMigrationCmd)
}

func handle(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		Print("red", "Provide the database and migration name. Example: migration make:migration yourdb create_users_table")
		return
	}

	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	createMigrationsFolder()
	createDatabaseFolder(args[0])
	createMigrationFile(args[0], args[1], timestamp)
	createRollbackFile(args[0], args[1], timestamp)
}

func createMigrationsFolder() {
	_, err := os.Stat("./migrations")

	if !os.IsNotExist(err) {
		if Verbose {
			Print("yellow", "Migrations folder already exists!")
		}
		return
	}

	if Verbose {
		Print("gray", "Creating migrations folder...")
	}

	os.Mkdir("migrations", 0755)

	Print("green", "Migrations folder created successfully!")
}

func createDatabaseFolder(database string) {
	_, err := os.Stat("./migrations/" + database)

	if !os.IsNotExist(err) {
		if Verbose {
			Print("yellow", "Database folder already exists!")
		}
		return
	}

	if Verbose {
		Print("gray", "Creating database folder...")
	}

	os.Mkdir("migrations/"+database, 0755)

	Print("green", "Database folder created successfully!")
}

func createMigrationFile(database string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + database + "/" + strings.Join(splittedPath, "/") + ".sql")
	defer file.Close()
	if err != nil {
		Print("red", "Error creating migration file")
		if Verbose {
			Print("gray", err.Error())
		}
		return
	}

	Print("green", "Migration file created successfully!")
}

func createRollbackFile(database string, path string, timestamp string) {
	splittedPath := strings.Split(path, "/")
	splittedPath[len(splittedPath)-1] = timestamp + "_" + splittedPath[len(splittedPath)-1]

	file, err := os.Create("migrations/" + database + "/" + strings.Join(splittedPath, "/") + "_rollback.sql")
	defer file.Close()
	if err != nil {
		Print("red", "Error creating rollback file")
		if Verbose {
			Print("gray", err.Error())
		}
		return
	}

	Print("green", "Rollback file created successfully!")
}
