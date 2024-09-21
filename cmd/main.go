package main

import (
	"context"
	"fmt"

	"github.com/suritasolutions/go-migration/internal/commands"
	"github.com/suritasolutions/go-migration/internal/db"
	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
)

func main() {
	rootCmd := commands.NewRootCommand()
	ctx := rootCmd.Context()

	var verboseFlag bool
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "verbose output")

	ctx = context.WithValue(ctx, "verbose", verboseFlag)
	logger := util.NewConsoleLogger()
	migrationFile := migration.NewMigrationFile(ctx, logger)

	rootCmd.AddCommand(commands.NewMigrateCommand(ctx, db.NewPostgresDB(ctx), migrationFile, logger))
	rootCmd.AddCommand(commands.NewMakeMigrationCommand(ctx, migrationFile, logger))

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
