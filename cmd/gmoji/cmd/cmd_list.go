package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runList(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(false, false, "")
	if err != nil {
		return err
	}

	if err := c.List(); err != nil {
		return err
	}

	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the list of gmojis",
	Long:  "Show the list of gmojis.",
	RunE:  runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
