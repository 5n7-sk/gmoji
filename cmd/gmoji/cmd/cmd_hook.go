package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

// HookOptions represents the options for hook command.
type HookOptions struct {
	Remove bool
}

var (
	hookOptions HookOptions
)

func runHook(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI()
	if err != nil {
		return err
	}

	if hookOptions.Remove {
		if err := c.RemoveHook(); err != nil {
			return err
		}
		return nil
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
	hookCmd.Flags().BoolVarP(&hookOptions.Remove, "remove", "r", false, "remove the commit hook")

	rootCmd.AddCommand(hookCmd)
}
