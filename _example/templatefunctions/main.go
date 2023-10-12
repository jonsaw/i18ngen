package main

import (
	"fmt"
	"html/template"
	"strings"

	i18ngen "github.com/jonsaw/i18ngen/_example/templatefunctions/i18n/generated"
)

func main() {
	templates := template.FuncMap{
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
	}

	fmt.Println(i18ngen.Load("en", templates).Welcome("John Doe"))
	fmt.Println(i18ngen.Load("ms", templates).Welcome("John Doe"))
}
