package main

import (
	"encoding/json"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func baseTemplate() string {
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

type {{.Name}} interface {
{{- range $jsonName, $val := .Translations}}
	{{- with (GetDefinition $definitions $jsonName) }}
		{{- if .Description}}
	// {{.Description}}
	//
	// In {{$lang}}, this message translates to: {{$val}}
		{{- else}}
	// No description provided for @{{$jsonName}}
	//
	// In {{$lang}}, this message translates to: {{$val}}
		{{- end}}
		{{- with .Placeholders}}
	{{(Title (CamelCase $jsonName))}}({{(PlaceholderTypes .)}}) string
		{{- else}}
	{{(Title (CamelCase $jsonName))}}() string
		{{- end}}
	{{- else}}
	// No definition provided for @{{$jsonName}}
	//
	// In {{$lang}}, this message translates to: {{$val}}
	{{(Title (CamelCase $jsonName))}}() string
	{{- end}}
{{- end }}
}

// Load loads the translation for the given language.
// Defaults to base language if not found.
func Load(lang string, funcMap template.FuncMap) {{.Name}} {
	switch lang {
	{{- range $lang := .Langs}}
	case "{{$lang}}":
		return &Translation{{(Title (CamelCase $lang))}}{TemplateFuncMap: funcMap}
	{{- end}}
	}
	return &Translation{{(Title (CamelCase .Lang))}}{TemplateFuncMap: funcMap}
}

// LoadLangMap loads the translation map for the given language.
// Typlically used to dynamically load translations by key.
// Defaults to base language if not found.
func LoadLangMap(lang string) LangMap {
	switch lang {
	{{- range $lang := .Langs}}
	case "{{$lang}}":
		return Translation{{(Title (CamelCase $lang))}}Map
	{{- end}}
	}
	return Translation{{(Title (CamelCase .Lang))}}Map
}

type LangMap map[string]string

func (m LangMap) Get(key string) (string, error) {
	if val, ok := m[key]; ok {
		return val, nil
	}

	return "", fmt.Errorf("key %q not found", key)
}

func (m LangMap) MustGet(key string) string {
	val, err := m.Get(key)
	if err != nil {
		return key
	}

	return val
}

func (m LangMap) GetTemplated(key string, values interface{}, templateFunc template.FuncMap) (string, error) {
	val, err := m.Get(key)
	if err != nil {
		return "", err
	}

	t := template.New("")
	if templateFunc != nil {
		t = t.Funcs(templateFunc)
	}
	t, err = t.Parse(val)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	err = t.Execute(&b, values)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (m LangMap) MustGetTemplated(key string, values interface{}, templateFunc template.FuncMap) string {
	val, err := m.GetTemplated(key, values, templateFunc)
	if err != nil {
		return key
	}

	return val
}
`
}

func generateBase(config Config, langs []string) error {
	baseLang := strings.Split(filepath.Base(config.Base), ".")[0]

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

	imports := []string{
		"html/template",
		"strings",
		"fmt",
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

	data := struct {
		Imports      []string
		Name         string
		Lang         string
		Langs        []string
		Translations map[string]string
		Definitions  map[string]Definition
		Package      string
	}{
		Imports:      imports,
		Name:         "Base",
		Lang:         baseLang,
		Langs:        langs,
		Translations: baseTranslations,
		Definitions:  definitions,
		Package:      config.Exec.Package,
	}

	tpl, err := template.New("base.tpl").Funcs(templateFns()).Parse(baseTemplate())
	if err != nil {
		return err
	}

	err = os.MkdirAll(config.Exec.Path, 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(config.Exec.Path, "/base.go"))
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
