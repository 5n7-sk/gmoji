package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runCopy(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI()
	if err != nil {
		return err
	}

	if err := c.RunCopy(); err != nil {
		return err
	}

	fmt.Printf("%s\n", color.GreenString("Copied to the clipboard successfully"))

	return nil
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy the commit message to the clipboard",
	Long:  "Copy the commit message to the clipboard.",
	RunE:  runCopy,
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
