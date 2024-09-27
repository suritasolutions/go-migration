package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
)

func NewMigrateCommand(
	ctx context.Context,
	logger util.Logger,
	migration migration.Migration,
) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run all migrations",
		Long:  "Run all migrations to the database.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logger.Fatal("Please provide the migration folder name.")
				return
			}

			if len(args) < 2 {
				logger.Fatal("Please provide a database name.")
				return
			}

			ctx = context.WithValue(ctx, "folder", args[0])
			ctx = context.WithValue(ctx, "database", args[1])

			migration.Migrate()
		},
	}

	migrateCmd.SetUsageTemplate(`Usage:
	migration migrate [database]

	`)
	migrateCmd.Example = `migration migrate [migration_folder] [database]`

	return migrateCmd
}
