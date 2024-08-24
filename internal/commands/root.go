package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "migration",
	Short:            "DB migration tool",
	Long:             "Database migration tool",
	TraverseChildren: true,
}

var Colors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"gray":   "\033[90m",
	"reset":  "\033[0m",
}

func Print(color string, message string) {
	fmt.Println(Colors[color], message, Colors["reset"])
}

var Verbose bool

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
