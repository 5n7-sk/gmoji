package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

// RootOptions represents the options for root command.
type RootOptions struct {
	Hook string
}

var (
	rootOptions RootOptions
)

func runRoot(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(false, rootOptions.Hook)
	if err != nil {
		return err
	}

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gmoji",
	Short: "gmoji is the Go Implementation of gitmoji",
	Long:  "gmoji is the Go Implementation of gitmoji.",
	RunE:  runRoot,
}

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&rootOptions.Hook, "hook", "", "hook path (.git/COMMIT_EDITMSG)")
}
