package commands

import (
	"context"
	"testing"

	"github.com/suritasolutions/go-migration/internal/migration"
	"github.com/suritasolutions/go-migration/util"
	"go.uber.org/mock/gomock"
)

func TestMigrationCommand(t *testing.T) {
	t.Run("Should fail if migration folder name param was not provided", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		mocker := gomock.NewController(t)

		migrationMock := migration.NewMockMigration(mocker)
		migrationMock.EXPECT().Migrate().Times(0)

		loggerMock := util.NewMockLogger(mocker)
		loggerMock.EXPECT().Fatal("Please provide the migration folder name.").Times(1)

		cmd := NewMigrateCommand(ctx, loggerMock, migrationMock)
		args := []string{}

		cmd.Run(cmd, args)
	})

	t.Run("Should fail if migration database name param was not provided", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		mocker := gomock.NewController(t)

		migrationMock := migration.NewMockMigration(mocker)
		migrationMock.EXPECT().Migrate().Times(0)

		loggerMock := util.NewMockLogger(mocker)
		loggerMock.EXPECT().Fatal("Please provide a database name.").Times(1)

		cmd := NewMigrateCommand(ctx, loggerMock, migrationMock)
		args := []string{"foldertest"}

		cmd.Run(cmd, args)
	})

	t.Run("Should start migration", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		mocker := gomock.NewController(t)

		migrationMock := migration.NewMockMigration(mocker)
		migrationMock.EXPECT().Migrate().Times(1)

		loggerMock := util.NewMockLogger(mocker)
		loggerMock.EXPECT().Fatal("").Times(0)

		cmd := NewMigrateCommand(ctx, loggerMock, migrationMock)
		args := []string{"foldertest", "database"}

		cmd.Run(cmd, args)
	})
}
