package cli

import (
	"fmt"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/skmatz/gmoji"
)

// Run runs CLI.
func (c CLI) Run(hookPath string) error {
	_, err := c.GitRoot()
	if err != nil {
		return err
	}

	p, err := c.ListPath()
	if err != nil {
		return err
	}

	gmojis, err := gmoji.NewGmojis(p)
	if err != nil {
		return err
	}

	answers, err := ask(gmojis)
	if err != nil {
		return err
	}

	commitMessage := c.commitMessage(gmojis[answers.Emoji].Code, answers.Title, answers.Message)

	if err := ioutil.WriteFile(hookPath, []byte(commitMessage), 0755); err != nil {
		return fmt.Errorf("%s not found (maybe nothing is staged)", hookPath)
	}

	return nil
}

// RunCopy runs CLI and copy the commit message to the clipboard.
func (c CLI) RunCopy() error {
	p, err := c.ListPath()
	if err != nil {
		return err
	}

	gmojis, err := gmoji.NewGmojis(p)
	if err != nil {
		return err
	}

	answers, err := ask(gmojis)
	if err != nil {
		return err
	}

	commitMessage := c.commitMessage(gmojis[answers.Emoji].Code, answers.Title, answers.Message)

	if err := clipboard.WriteAll(commitMessage); err != nil {
		return err
	}

	return nil
}
