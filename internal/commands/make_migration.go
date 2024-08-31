package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
)

var MakeMigrationCmd = &cobra.Command{
	Use:   "make:migration",
	Short: "Generate a new migration",
	Long:  `Generate a new migration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			util.Print("red", "Provide the migration folder and file name. Example: migration make:migration [migration_folder_name] [migration_file_name]")
			return
		}

		migrationFile := migration.NewMigrationFile(
			context.WithValue(cmd.Context(), "verbose", Verbose),
		)

		migrationFile.Create(args[0], args[1])
	},
}

func init() {
	MakeMigrationCmd.Example = "migration make:migration [migration_folder_name] [migration_file_name]"
	rootCmd.AddCommand(MakeMigrationCmd)
}
