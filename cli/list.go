package cli

import (
	"fmt"

	"github.com/skmatz/gmoji"
)

// List prints the list of gmojis.
func (c CLI) List() error {
	p, err := c.ListPath()
	if err != nil {
		return err
	}

	gmojis, err := gmoji.NewGmojis(p)
	if err != nil {
		return err
	}

	for _, s := range gmojis.Selection() {
		fmt.Println(s)
	}

	return nil
}
