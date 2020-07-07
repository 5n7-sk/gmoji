package cli

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/skmatz/gmoji"
)

type answers struct {
	Emoji   int    // index
	Title   string // commit title
	Message string // commit message
}

func ask(gmojis gmoji.Gmojis) (*answers, error) {
	q := []*survey.Question{
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

	var a answers

	if err := survey.Ask(q, &a); err != nil {
		return nil, err
	}

	return &a, nil
}
