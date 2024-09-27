package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
)

func NewMakeMigrationCommand(
	ctx context.Context,
	migrationFile migration.MigrationFile,
	logger util.Logger,
) *cobra.Command {
	makeMigrationCmd := &cobra.Command{
		Use:   "make:migration",
		Short: "Generate a new migration",
		Long:  `Generate a new migration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				logger.Fatal("Provide the migration folder and file name. Example: migration make:migration [migration_folder_name] [migration_file_name]")
				return
			}

			migrationFile.Create(args[0], args[1])
		},
	}

	makeMigrationCmd.Example = "migration make:migration [migration_folder_name] [migration_file_name]"

	return makeMigrationCmd
}
