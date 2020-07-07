package gmoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

const (
	// JSONPath is the path to the original gitmojis JSON file.
	JSONPath string = "https://raw.githubusercontent.com/carloscuesta/gitmoji/master/src/data/gitmojis.json"
)

// Gmoji represents a gmoji.
type Gmoji struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Name        string `json:"name"`
}

// Gmojis represents a list of Gmojis.
type Gmojis []Gmoji

// Gitmojis represents an original gitmojis.
type Gitmojis struct {
	Gitmojis Gmojis `json:"gitmojis"`
}

// NewGmojis returns a new Gmojis.
func NewGmojis(path string) (Gmojis, error) {
	g := &Gitmojis{}
	if err := g.readJSON(path); err != nil {
		return nil, err
	}

	return g.Gitmojis, nil
}

func (g *Gitmojis) readJSON(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s not found; download with 'gmoji init'", path)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &g); err != nil {
		return err
	}

	return nil
}

// Selection returns a formatted list of gmojis.
// FIXME: The width between the emoji and the hyphen is not consistent.
func (g Gmojis) Selection() []string {
	var gmojis []string
	for _, gmoji := range g {
		gmojis = append(gmojis, fmt.Sprintf("%s - %s - %s", gmoji.Emoji, color.BlueString(gmoji.Code), gmoji.Description))
	}

	return gmojis
}
