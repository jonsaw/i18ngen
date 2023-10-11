package main

import (
	"encoding/json"
	"html/template"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func templateFns() template.FuncMap {
	return template.FuncMap{
		"Title": func(v string) string {
			if v == "" {
				return ""
			}
			return strings.ToUpper(v[0:1]) + v[1:]
		},
		"CamelCase": func(v string) string {
			v = regexp.MustCompile("[^a-zA-Z0-9_ -]+").ReplaceAllString(v, "")
			v = strings.ReplaceAll(v, "_", " ")
			v = strings.ReplaceAll(v, "-", " ")
			v = cases.Title(language.AmericanEnglish, cases.NoLower).String(v)
			v = strings.ReplaceAll(v, " ", "")
			if len(v) > 0 {
				v = strings.ToLower(v[:1]) + v[1:]
			}
			return v
		},
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
		"StartsWith": func(v string, prefix string) bool {
			return strings.HasPrefix(v, prefix)
		},
		"GetDefinition": func(fields map[string]interface{}, key string) *Definition {
			v, ok := fields["@"+key]
			if !ok {
				return nil
			}

			b, err := json.Marshal(v)
			if err != nil {
				return nil
			}

			d := &Definition{}
			err = json.Unmarshal(b, d)
			if err != nil {
				return nil
			}
			return d
		},
		"GetTranslation": func(translations map[string]interface{}, key string) string {
			v, ok := translations[key]
			if !ok {
				return key
			}
			return v.(string)
		},
		"PlaceholderTypes": func(d map[string]Placeholder) string {
			var types []string
			for k, v := range d {
				types = append(types, k+" "+v.Type)
			}
			return strings.Join(types, ", ")
		},
	}
}
