package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Placeholder struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	Example string `json:"example"`
}

type Definition struct {
	Description  string        `json:"description"`
	Placeholders []Placeholder `json:"placeholders"`
}

func main() {
	app := &cli.App{
		Name: "i18ngen",
		Commands: []*cli.Command{
			generate(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generate() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "Generate translations",
		Action: func(c *cli.Context) error {
			config, err := readConfig("i18ngen.yml")
			if err != nil {
				return err
			}

			trs, langs, err := config.ListTranslationFiles()
			if err != nil {
				return err
			}

			err = generateBase(*config, langs)
			if err != nil {
				return err
			}

			for _, tr := range trs {
				err = generateLang(*config, tr)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
