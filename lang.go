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

import (
	"html/template"
	"strings"
)

{{- $lang := .Lang}}
{{- $fields := .Fields}}
{{- $translations := .Translations}}

{{- $name := print "Translation" (Title (CamelCase .Lang))}}

type {{$name}} struct {}

{{- range $jsonName, $val := .Fields}}
{{- if not (StartsWith $jsonName "@")}}
{{- with (GetDefinition $fields $jsonName) }}
{{- with .Placeholders}}
func (l *{{$name}}) {{(Title $jsonName)}}({{(PlaceholderTypes .)}}) string {
	t, err := template.New("").Parse("{{GetTranslation $translations $jsonName}}")
	if err != nil {
		return "{{$jsonName}}"
	}
	var b strings.Builder
	err = t.Execute(&b, map[string]interface{}{
		{{- range $placeholderName, $placeholder := .}}
		"{{$placeholderName}}": {{$placeholderName}},
		{{- end}}
	})
	if err != nil {
		return "{{$jsonName}}"
	}
	return b.String()
}
{{- else}}
func (l *{{$name}}) {{(Title $jsonName)}}() string {
	return "{{GetTranslation $translations $jsonName}}"
}
{{- end}}
{{- else}}
func (l *{{$name}}) {{(Title $jsonName)}}() string {
	return "{{GetTranslation $translations $jsonName}}"
}
{{- end}}
{{- end}}
{{- end }}
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

	var fields map[string]interface{}
	err = json.Unmarshal(b, &fields)
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

	var translations map[string]interface{}
	err = json.Unmarshal(b, &translations)
	if err != nil {
		return err
	}

	data := struct {
		Lang         string
		Fields       map[string]interface{}
		Translations map[string]interface{}
		Package      string
	}{
		lang,
		fields,
		translations,
		config.Exec.Package,
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
