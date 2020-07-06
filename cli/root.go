package cli

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/skmatz/gmoji"
)

// Run runs CLI.
func (c CLI) Run() error {
	if c.CheckGit {
		_, err := c.GitRoot()
		if err != nil {
			return err
		}
	}

	p, err := c.ListPath()
	if err != nil {
		return err
	}

	gmojis, err := gmoji.NewGmojis(p)
	if err != nil {
		return err
	}

	// FIXME: The width between the emoji and the hyphen is not consistent.
	questions := []*survey.Question{
		{
			Name: "emoji",
			Prompt: &survey.Select{
				Message:  "Choose a gmoji",
				Options:  gmojis.Selection(),
				PageSize: 10,
				Filter: func(filter string, value string, index int) bool {
					f := strings.ToLower(filter)
					g := gmojis[index]
					return fuzzy.Match(f, strings.ToLower(g.Description)) || fuzzy.Match(f, strings.ToLower(g.Name))
				},
			},
			Validate: survey.Required,
		},
		{
			Name:     "title",
			Prompt:   &survey.Input{Message: "Enter the commit title"},
			Validate: survey.Required,
		},
		{
			Name:   "message",
			Prompt: &survey.Input{Message: "Enter the commit message"},
		},
	}

	answers := struct {
		Emoji   int    // index
		Title   string // commit title
		Message string // commit message
	}{}

	if err := survey.Ask(questions, &answers); err != nil {
		return err
	}

	commitMessage := c.commitMessage(gmojis[answers.Emoji].Code, answers.Title, answers.Message)

	switch c.Clipboard {
	case true:
		if err := clipboard.WriteAll(commitMessage); err != nil {
			return err
		}
	case false:
		if err := ioutil.WriteFile(c.HookPath, []byte(commitMessage), 0755); err != nil {
			return fmt.Errorf("%s not found (maybe nothing is staged)", c.HookPath)
		}
	}

	return nil
}
