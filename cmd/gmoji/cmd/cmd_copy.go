package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runCopy(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(true, "")
	if err != nil {
		return err
	}

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy commit message to clipboard",
	Long:  "Copy commit message to clipboard.",
	RunE:  runCopy,
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
