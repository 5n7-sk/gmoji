package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the gmoji version.
	Version = "unset"
)

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("gmoji v%s\n", Version)

	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version.",
	RunE:  runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
