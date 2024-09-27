package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:              "migration",
		Short:            "DB migration tool",
		Long:             "Database migration tool",
		TraverseChildren: true,
	}
}
