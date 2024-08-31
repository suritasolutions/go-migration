package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/suritasolutions/go-migration/internal/db"
	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run all migrations",
	Long:  "Run all migrations to the database.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			util.Print("red", "Please provide the migration folder name.")
			return
		}

		if len(args) < 2 {
			util.Print("red", "Please provide a database name.")
			return
		}

		ctx := context.WithValue(cmd.Context(), "verbose", Verbose)
		ctx = context.WithValue(ctx, "folder", args[0])
		ctx = context.WithValue(ctx, "database", args[1])

		migration := migration.NewMigration(
			ctx,
			db.NewPostgresDB(ctx),
			migration.NewMigrationFile(ctx),
		)

		migration.Migrate()
	},
}

func init() {
	MigrateCmd.SetUsageTemplate(`Usage:
	migration migrate [database]

	`)
	MigrateCmd.Example = `migration migrate [migration_folder] [database]`
	rootCmd.AddCommand(MigrateCmd)
}
