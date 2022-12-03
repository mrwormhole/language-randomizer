package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"gopkg.in/yaml.v3"
)

//go:embed languages/languages.yaml
var raw []byte

type language struct {
	LanguageID int `yaml:"language_id"`
	Name       string
	Type       string   `yaml:"type"`
	Color      string   `yaml:"color"`
	Group      string   `yaml:"group"`
	Aliases    []string `yaml:"aliases"`
	Extensions []string `yaml:"extensions"`
}

func main() {
	amount := flag.Int("select", 1, "languages amount you would like to select")
	flag.Parse()

	var languages map[string]language
	err := yaml.Unmarshal(raw, &languages)
	if err != nil {
		panic(err)
	}

	for k, v := range languages {
		if v.Type != "programming" || v.Group != "" || v.Color == "" || len(v.Extensions) == 0 {
			delete(languages, k)
			continue
		}
		v.Name = k
		languages[k] = v
	}

	for i := *amount; i > 0; i-- {
		selectedLanguage := pick(languages)
		if selectedLanguage != nil {
			display(selectedLanguage)
		}
	}
}

func pick(m map[string]language) *language {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	val := r.Intn(len(m))
	for _, v := range m {
		if val == 0 {
			return &v
		}
		val--
	}
	return nil
}

func display(l *language) {
	fmt.Println("-----------------------------------------")
	rgb, _ := putils.RGBFromHEX(l.Color)
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithRGB(l.Name, rgb)).Render()

	if len(l.Aliases) > 0 {
		pterm.Info.Println("🌚 also called as: ", l.Aliases)
	}
	pterm.Info.Println("⭐ file extensions: ", l.Extensions)
	fmt.Println("-----------------------------------------")
}
