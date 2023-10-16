package main

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
	config, err := readConfig("i18ngen.yml")
	if err != nil {
		panic(err)
	}

	trs, langs, err := config.ListTranslationFiles()
	if err != nil {
		panic(err)
	}

	err = generateBase(*config, langs)
	if err != nil {
		panic(err)
	}

	for _, tr := range trs {
		err = generateLang(*config, tr)
		if err != nil {
			panic(err)
		}
	}
}
