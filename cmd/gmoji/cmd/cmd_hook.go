package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runHook(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(true, false, "")
	if err != nil {
		return err
	}

	if err := c.Hook(); err != nil {
		return err
	}

	return nil
}

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Set the commit hook",
	Long:  "Set the commit hook.",
	RunE:  runHook,
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
