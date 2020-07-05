package cmd

import (
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runInit(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(false, false, "")
	if err != nil {
		return err
	}

	if err := c.Init(); err != nil {
		return err
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Download the list of gmojis",
	Long:  "Download the list of gmojis.",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
