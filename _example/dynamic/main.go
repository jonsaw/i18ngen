package main

import (
	"fmt"

	i18ngen "github.com/jonsaw/i18ngen/_example/dynamic/i18n/generated"
)

func main() {
	type customError struct {
		key      string
		label    string
		settings []string
	}

	errors := []customError{
		{"invalidAlreadyInUse", "email", nil},
		{"invalidAlreadyRegistered", "email", nil},
		{"invalidGreaterThan", "password", []string{"8"}},
		{"invalidMinLength", "password", []string{"8"}},
		{"invalidRequired", "email", nil},
	}

	lang := i18ngen.LoadLangMap("en")
	for _, cErr := range errors {
		label, err := lang.Get(cErr.label)
		if err != nil {
			panic(err)
		}
		message, err := lang.GetTemplated(cErr.key, map[string]interface{}{
			"label":    label,
			"settings": cErr.settings,
		}, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}

	lang = i18ngen.LoadLangMap("ms")
	for _, cErr := range errors {
		label, err := lang.Get(cErr.label)
		if err != nil {
			panic(err)
		}
		message, err := lang.GetTemplated(cErr.key, map[string]interface{}{
			"label":    label,
			"settings": cErr.settings,
		}, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
}
