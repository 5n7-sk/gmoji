package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/skmatz/gmoji/cli"
	"github.com/spf13/cobra"
)

func runCopy(cmd *cobra.Command, args []string) error {
	c, err := cli.NewCLI(false, true, "")
	if err != nil {
		return err
	}

	if err := c.Run(); err != nil {
		return err
	}

	fmt.Printf("%s\n", color.GreenString("Copied to clipboard successfully"))

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
