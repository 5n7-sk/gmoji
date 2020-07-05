package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/skmatz/gmoji"
)

// Init downloads the list of gmojis.
func (c CLI) Init() error {
	p, err := c.ListPath()
	if err != nil {
		return err
	}

	if err := c.Wget(gmoji.JSONPath, p); err != nil {
		return err
	}

	fmt.Printf("%s -> %s\n", color.GreenString(gmoji.JSONPath), color.GreenString(p))

	return nil
}
