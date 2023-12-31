package main

import (
	"encoding/json"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func langTemplate() string {
	return `package {{.Package}}

// This code generated by go generate.
// DO NOT EDIT BY HAND!

{{- with .Imports}}
{{- if gt (len .) 1}}
import (
{{- range .}}
	"{{.}}"
{{- end}}
)
{{- else if gt (len .) 0}}

import "{{(index . 0)}}"
{{- end}}
{{- end}}

{{- $lang := .Lang}}
{{- $definitions := .Definitions}}
{{- $translations := .Translations}}

{{- $name := print "Translation" (Title (CamelCase .Lang))}}

type {{$name}} struct {
	TemplateFuncMap template.FuncMap
}

{{- range $jsonName, $val := .BaseTranslations}}
{{- with (GetDefinition $definitions $jsonName) }}
{{- with .Placeholders}}
func (l *{{$name}}) {{(Title (CamelCase $jsonName))}}({{(PlaceholderTypes .)}}) string {
	return Translation{{(Title (CamelCase $lang))}}Map.MustGetTemplated(
		"{{$jsonName}}",
		map[string]interface{}{
			{{- range $idx, $placeholder := .}}
			"{{$placeholder.Label}}": {{CamelCase $placeholder.Label}},
			{{- end}}
		},
		l.TemplateFuncMap,
	)
}
{{- else}}
func (l *{{$name}}) {{(Title (CamelCase $jsonName))}}() string {
	return Translation{{(Title (CamelCase $lang))}}Map.MustGet("{{$jsonName}}")
}
{{- end}}
{{- else}}
func (l *{{$name}}) {{(Title (CamelCase $jsonName))}}() string {
	return Translation{{(Title (CamelCase $lang))}}Map.MustGet("{{$jsonName}}")
}
{{- end}}
{{- end }}

var Translation{{(Title (CamelCase $lang))}}Map = LangMap{
	{{- range $jsonName, $val := .BaseTranslations}}
	"{{$jsonName}}": "{{GetTranslation $translations $jsonName}}",
	{{- end}}
}
`
}

func generateLang(config Config, langFilepath string) error {
	lang := strings.Split(filepath.Base(langFilepath), ".")[0]

	in, err := os.Open(config.Base)
	if err != nil {
		return err
	}
	defer in.Close()

	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	var rawBaseDefinitions map[string]interface{}
	err = json.Unmarshal(b, &rawBaseDefinitions)
	if err != nil {
		return err
	}

	in, err = os.Open(langFilepath)
	if err != nil {
		return err
	}
	defer in.Close()

	b, err = io.ReadAll(in)
	if err != nil {
		return err
	}

	var rawTranslations map[string]interface{}
	err = json.Unmarshal(b, &rawTranslations)
	if err != nil {
		return err
	}

	imports := []string{
		"html/template",
	}

	baseTranslations := map[string]string{}
	definitions := map[string]Definition{}
	for k, f := range rawBaseDefinitions {
		if f == nil {
			continue
		}

		if !strings.HasPrefix(k, "@") {
			if sf, ok := f.(string); ok {
				baseTranslations[k] = sf
			}
		}
		if strings.HasPrefix(k, "@") {
			b, err := json.Marshal(f)
			if err != nil {
				return err
			}

			d := Definition{}
			err = json.Unmarshal(b, &d)
			if err != nil {
				return err
			}

			definitions[k] = d
		}
	}

	translations := map[string]string{}
	for k, f := range rawTranslations {
		if f == nil {
			continue
		}
		if !strings.HasPrefix(k, "@") {
			if sf, ok := f.(string); ok {
				translations[k] = sf
			}
		}
	}

	data := struct {
		Imports          []string
		Lang             string
		BaseTranslations map[string]string
		Translations     map[string]string
		Definitions      map[string]Definition
		Package          string
	}{
		Imports:          uniqueSliceElements(imports),
		Lang:             lang,
		BaseTranslations: baseTranslations,
		Translations:     translations,
		Definitions:      definitions,
		Package:          config.Exec.Package,
	}

	tpl, err := template.New("lang.tpl").Funcs(templateFns()).Parse(langTemplate())
	if err != nil {
		return err
	}

	err = os.MkdirAll(config.Exec.Path, 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(config.Exec.Path, lang+".go"))
	if err != nil {
		return err
	}
	defer out.Close()

	err = tpl.Execute(out, data)
	if err != nil {
		return err
	}

	return nil
}
