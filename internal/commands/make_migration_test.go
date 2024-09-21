package commands

import (
	"context"
	"testing"

	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
	"go.uber.org/mock/gomock"
)

func TestMakeMigrationCommand(t *testing.T) {

	// Test if the command fails when no arguments are provided
	t.Run("Should return error if run with no args", func(t *testing.T) {
		ctx := context.Background()

		mocker := gomock.NewController(t)
		defer mocker.Finish()

		migrationFileMock := migration.NewMockMigrationFile(mocker)
		migrationFileMock.EXPECT().Create("", "").Times(0)

		loggerMock := util.NewMockLogger(mocker)
		loggerMock.EXPECT().Fatal(
			"Provide the migration folder and file name. Example: migration make:migration [migration_folder_name] [migration_file_name]",
		).Times(1)

		cmd := NewMakeMigrationCommand(ctx, migrationFileMock, loggerMock)
		args := []string{}

		cmd.Run(cmd, args)
	})

	t.Run("Should create a new migration file", func(t *testing.T) {
		ctx := context.Background()

		mocker := gomock.NewController(t)

		migrationFileMock := migration.NewMockMigrationFile(mocker)
		migrationFileMock.EXPECT().Create("db", "path").Times(1)

		loggerMock := util.NewMockLogger(mocker)

		cmd := NewMakeMigrationCommand(ctx, migrationFileMock, loggerMock)
		args := []string{"db", "path"}

		cmd.Run(cmd, args)
	})
}
