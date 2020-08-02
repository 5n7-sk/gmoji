package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// Hook sets the commit hook.
func (c CLI) Hook() error {
	p, err := c.hookPath()
	if err != nil {
		return err
	}

	contents := `#!/bin/sh
exec < /dev/tty
gmoji --hook $1
`
	if err := ioutil.WriteFile(p, []byte(contents), 0755); err != nil {
		return err
	}

	fmt.Printf("%s\n", color.GreenString("gmoji commit hook created successfully"))

	return nil
}

// RemoveHook removes the commit hook.
func (c CLI) RemoveHook() error {
	p, err := c.hookPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", p)
	}

	if err := os.Remove(p); err != nil {
		return err
	}

	fmt.Printf("%s\n", color.GreenString("gmoji commit hook removed successfully"))

	return nil
}

func (c CLI) hookPath() (string, error) {
	gitRoot, err := c.GitRoot()
	if err != nil {
		return "", err
	}

	p := filepath.Join(gitRoot, ".git", "hooks", "prepare-commit-msg")

	return p, nil
}
