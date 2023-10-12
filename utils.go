package main

import (
	"html/template"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func uniqueSliceElements[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}

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
		"GetDefinition": func(fields map[string]Definition, key string) *Definition {
			d, ok := fields["@"+key]
			if !ok {
				return nil
			}
			return &d
		},
		"GetTranslation": func(translations map[string]string, key string) string {
			v, ok := translations[key]
			if !ok {
				return key
			}
			return v
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
